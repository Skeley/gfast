package wechat

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/errors/gcode"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gctx"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/tiger1103/gfast/v3/internal/app/shenaijia/model"
	"github.com/tiger1103/gfast/v3/internal/app/shenaijia/service"
)

func init() {
	service.RegisterWeChat(New())
}

func New() *sWeChat {
	s := &sWeChat{}
	s.init()
	return s
}

type sWeChat struct {
	apps map[string]*sApp
}

func (s *sWeChat) init() {
	s.apps = make(map[string]*sApp)
	s.apps[s.readerAppType()] = &sApp{}
	s.apps[s.readerAppType()].init("wechat.reader")

	s.apps[s.writerAppType()] = &sApp{}
	s.apps[s.writerAppType()].init("wechat.writer")
}

func (s *sWeChat) userType2AppType(userType uint) (string, error) {
	switch userType {
	case 1, 3:
		return s.readerAppType(), nil
	case 2:
		return s.writerAppType(), nil
	default:
		return "", gerror.NewCodef(gcode.CodeInvalidParameter, `invalid user type %d`, userType)
	}
}

func (s *sWeChat) readerAppType() string {
	return "reader"
}

func (s *sWeChat) writerAppType() string {
	return "writer"
}

func (s *sWeChat) Jscode2Session(ctx context.Context, userType uint, loginCode string) (result *model.Jscode2SessionResp, err error) {
	appType, e := s.userType2AppType(userType)
	if e != nil {
		return nil, e
	}
	return s.apps[appType].Jscode2Session(ctx, loginCode)
}

func (s *sWeChat) GetPhoneNumber(ctx context.Context, userType uint, code string) (string, error) {
	appType, e := s.userType2AppType(userType)
	if e != nil {
		return "", e
	}
	return s.apps[appType].GetPhoneNumber(ctx, code)
}

type sApp struct {
	appid  string
	secret string
	token  *model.TokenRsp

	tokenRefreshTicker *time.Ticker
	httpCli            *http.Client
}

func (s *sApp) Jscode2Session(ctx context.Context, loginCode string) (result *model.Jscode2SessionResp, err error) {
	result = &model.Jscode2SessionResp{}

	baseUrl := "https://api.weixin.qq.com/sns/jscode2session"
	params := url.Values{}
	params.Add("appid", s.appid)
	params.Add("secret", s.secret)
	params.Add("js_code", loginCode)
	params.Add("grant_type", "authorization_code")

	fullURL := baseUrl + "?" + params.Encode()
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, fullURL, nil)
	resp, e := s.httpCli.Do(req)
	if e != nil {
		return nil, gerror.Wrap(e, "wechat session")
	}
	defer resp.Body.Close()
	body, e := io.ReadAll(resp.Body)
	if e != nil {
		return nil, gerror.Wrap(e, "wechat session")
	}
	_ = json.Unmarshal(body, &result)

	if result.ErrCode != 0 {
		return nil, gerror.Newf("微信接口返回错误，错误码: %d，错误信息: %s", result.ErrCode, result.ErrMsg)
	}
	return
}

func (s *sApp) GetPhoneNumber(ctx context.Context, code string) (string, error) {
	baseUrl := "https://api.weixin.qq.com/wxa/business/getuserphonenumber"
	params := url.Values{}
	params.Add("access_token", s.token.AccessToken)
	fullURL := baseUrl + "?" + params.Encode()

	data, _ := json.Marshal(model.GetPhoneNumberReq{
		// OpenId: openid,
		Code: code,
	})

	req, _ := http.NewRequestWithContext(ctx, "POST", fullURL, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	rsp, e := s.httpCli.Do(req)
	if e != nil {
		return "", gerror.Wrap(e, "wechat get phone number")
	}
	defer rsp.Body.Close()
	body, e := io.ReadAll(rsp.Body)
	if e != nil {
		return "", gerror.Wrap(e, "refresh wechat access token")
	}
	phoneNumberRsp := &model.GetPhoneNumberRsp{}
	_ = json.Unmarshal(body, phoneNumberRsp)
	if phoneNumberRsp.ErrCode != 0 {
		return "", gerror.Newf("wechat get phone number, code: %d, msg: %s",
			phoneNumberRsp.ErrCode, phoneNumberRsp.ErrMsg)
	}
	return phoneNumberRsp.PhoneInfo.PhoneNumber, nil
}

func (s *sApp) refreshToken() error {
	baseUrl := "https://api.weixin.qq.com/cgi-bin/stable_token"
	data, _ := json.Marshal(model.TokenReq{
		Appid:     s.appid,
		Secret:    s.secret,
		GrantType: "client_credential",
	})
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	req, _ := http.NewRequestWithContext(ctx, "POST", baseUrl, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	rsp, e := s.httpCli.Do(req)
	if e != nil {
		return gerror.Wrap(e, "refresh wechat access token")
	}
	defer rsp.Body.Close()
	body, e := io.ReadAll(rsp.Body)
	if e != nil {
		return gerror.Wrap(e, "refresh wechat access token")
	}

	tokenRsp := &model.TokenRsp{}
	json.Unmarshal(body, tokenRsp)
	s.token = tokenRsp
	return nil
}

func (s *sApp) tokenTTL() time.Duration {
	return time.Duration(s.token.ExpiresIn)*time.Second - time.Minute*5
}

func (s *sApp) init(confPath string) {
	rawAppConf, err := gcfg.Instance().Get(gctx.GetInitCtx(), confPath)
	if err != nil {
		g.Throw("wechat init fail, invalid app conf")
	}
	conf := rawAppConf.MapStrStr()
	appid, exist := conf["appid"]
	if !exist || len(appid) == 0 {
		g.Throw("wechat init fail, invalid appid")
	}
	secret, exist := conf["secret"]
	if !exist || len(secret) == 0 {
		g.Throw("wechat init fail, invalid secret")
	}
	s.appid = appid
	s.secret = secret
	s.httpCli = &http.Client{
		Timeout: time.Second * 2,
	}

	if e := s.refreshToken(); e != nil {
		g.Throw("wechat init fail, refresh token")
	}

	s.tokenRefreshTicker = time.NewTicker(s.tokenTTL())
	go func() {
		for range s.tokenRefreshTicker.C {
			if e := s.refreshToken(); e != nil {
				g.Log().Errorf(gctx.GetInitCtx(), "refresh token fail: %v", e)
				s.tokenRefreshTicker.Reset(time.Second)
				continue
			}
			g.Log().Info(gctx.GetInitCtx(), "refresh token success")
			s.tokenRefreshTicker.Reset(s.tokenTTL())
		}
	}()
}

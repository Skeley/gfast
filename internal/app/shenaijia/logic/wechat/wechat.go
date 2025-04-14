package wechat

import (
	"bytes"
	"context"
	"encoding/json"
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
	appid  string
	secret string
	token  *model.TokenRsp

	tokenRefreshTicker *time.Ticker
	httpCli            *http.Client
}

func (s *sWeChat) Jscode2Session(ctx context.Context, loginCode string) (result *model.Jscode2SessionResp, err error) {
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

func (s *sWeChat) GetPhoneNumber(ctx context.Context, code string) (string, error) {
	baseUrl := "https://api.weixin.qq.com/cgi-bin/stable_token"
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

func (s *sWeChat) refreshToken() error {
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

func (s *sWeChat) tokenTTL() time.Duration {
	return time.Duration(s.token.ExpiresIn)*time.Second - time.Minute*5
}

func (s *sWeChat) init() {
	var ctx = gctx.New()
	appidValue, err := gcfg.Instance().Get(ctx, "wechat.appid")
	if err != nil || appidValue.String() == "" {
		g.Throw("wechat init fail, invalid appid")
	}
	secretValue, err := gcfg.Instance().Get(ctx, "wechat.secret")
	if err != nil || secretValue.String() == "" {
		g.Throw("wechat init fail, invalid secret")
	}
	s.appid = appidValue.String()
	s.secret = secretValue.String()
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

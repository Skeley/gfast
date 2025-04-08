package file

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path/filepath"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/google/uuid"
	"github.com/h2non/filetype"
	"github.com/tencentyun/cos-go-sdk-v5"

	v1 "github.com/tiger1103/gfast/v3/api/shenaijia/v1"
	"github.com/tiger1103/gfast/v3/internal/app/shenaijia/service"
	"github.com/tiger1103/gfast/v3/library/liberr"
)

func init() {
	service.RegisterFile(New())
}

func New() *sFile {
	s := &sFile{}
	s.cc = newCosCli()
	return s
}

type sFile struct {
	cc *cosCli
}

func (s *sFile) Upload(ctx context.Context, req *v1.UploadFileReq) (res *v1.UploadFileRes, err error) {
	res = &v1.UploadFileRes{}
	f, e := req.File.Open()
	if e != nil {
		err = gerror.Wrapf(err, `UploadFile.Open failed`)
		return res, err
	}
	defer f.Close()

	url, e := s.cc.UploadFile(ctx, req.File.Filename, f)
	liberr.ErrIsNil(ctx, e, "上传资源异常")
	res.URL = url
	return
}

type cosCli struct {
	cli    *cos.Client
	cosURL string
}

type cosConf struct {
	Bucket    string `json:"bucket"`
	Region    string `json:"region"`
	SecretID  string `json:"secret_id"`
	SecretKey string `json:"secret_key"`
}

func newCosConf() *cosConf {
	var ctx = gctx.New()

	bucket, _ := gcfg.Instance().Get(ctx, "cos.bucket")
	region, _ := gcfg.Instance().Get(ctx, "cos.region")
	secretId, _ := gcfg.Instance().Get(ctx, "cos.secret_id")
	secretKey, _ := gcfg.Instance().Get(ctx, "cos.secret_key")

	if bucket == nil || region == nil || secretId == nil || secretKey == nil {
		return nil
	}

	conf := &cosConf{
		Bucket:    bucket.String(),
		Region:    region.String(),
		SecretID:  secretId.String(),
		SecretKey: secretKey.String(),
	}
	g.Log().Infof(ctx, "cos.bucket: %s, cos.region: %s, cos.secret_id: %s, cos.secret_key: %s",
		conf.Bucket, conf.Region, conf.SecretID, conf.SecretKey)
	return conf
}

func newCosCli() *cosCli {
	cc := &cosCli{}

	conf := newCosConf()
	if conf == nil {
		g.Throw("cos init fail")
	}
	urlStr := fmt.Sprintf("https://%s.cos.%s.myqcloud.com", conf.Bucket, conf.Region)
	u, _ := url.Parse(urlStr)
	b := &cos.BaseURL{BucketURL: u}
	cc.cli = cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  conf.SecretID,
			SecretKey: conf.SecretKey,
		},
	})
	cc.cosURL = urlStr
	return cc
}

func detectFileExt(data []byte) (string, error) {
	kind, err := filetype.Match(data)
	if err != nil {
		return "", errors.New("不支持的文件类型")
	}
	if kind.MIME.Type != "image" && kind.MIME.Value != "application/pdf" {
		return "", errors.New("不支持的文件类型")
	}
	return kind.Extension, nil
}

func hasExtension(filename string) bool {
	ext := filepath.Ext(filename)
	return ext != ""
}

func (cc *cosCli) UploadFile(ctx context.Context, fileName string, reader io.Reader) (string, error) {
	header := make([]byte, 261)
	_, err := io.ReadFull(reader, header)
	if err != nil {
		return "", gerror.New("无效文件，文件大小必须大于261字节")
	}
	ext, err := detectFileExt(header)
	if err != nil {
		return "", err
	}
	prefix := "images/"
	if ext == "pdf" {
		prefix = "pdf/"
	}
	name := prefix + uuid.New().String()
	if !hasExtension(fileName) {
		fileName = name + "." + ext
	}
	_, err = cc.cli.Object.Put(ctx, name, io.MultiReader(bytes.NewReader(header), reader), nil)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf(cc.cosURL+"/%s", name), nil
}

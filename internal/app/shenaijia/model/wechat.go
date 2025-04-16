package model

type Jscode2SessionResp struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

type TokenReq struct {
	GrantType string `json:"grant_type"`
	Appid     string `json:"appid"`
	Secret    string `json:"secret"`
}

type TokenRsp struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

type GetPhoneNumberReq struct {
	OpenId string `json:"openid"`
	Code   string `json:"code"`
}

type PhoneInfo struct {
	PhoneNumber     string `json:"phoneNumber"`
	PurePhoneNumber string `json:"purePhoneNumber"`
	CountryCode     string `json:"countryCode"`
}

type GetPhoneNumberRsp struct {
	ErrCode   int        `json:"errcode"`
	ErrMsg    string     `json:"errmsg"`
	PhoneInfo *PhoneInfo `json:"phone_info"`
}

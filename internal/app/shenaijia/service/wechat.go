// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/tiger1103/gfast/v3/internal/app/shenaijia/model"
)

type (
	IWeChat interface {
		Jscode2Session(ctx context.Context, loginCode string) (result *model.Jscode2SessionResp, err error)
		GetPhoneNumber(ctx context.Context, code string) (string, error)
	}
)

var (
	localWeChat IWeChat
)

func WeChat() IWeChat {
	if localWeChat == nil {
		panic("implement not found for interface IWeChat, forgot register?")
	}
	return localWeChat
}

func RegisterWeChat(i IWeChat) {
	localWeChat = i
}

// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	v1 "github.com/tiger1103/gfast/v3/api/shenaijia/v1"
)

type (
	IUser interface {
		BoundCommunity(ctx context.Context, req *v1.UserBoundCommunityReq) (res *v1.UserBoundCommunityRes, err error)
		BindCommunity(ctx context.Context, req *v1.UserBindCommunityReq) (res *v1.UserBindCommunityRes, err error)
		UnbindCommunity(ctx context.Context, req *v1.UserUnbindCommunityReq) (res *v1.UserUnbindCommunityRes, err error)
	}
)

var (
	localUser IUser
)

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}

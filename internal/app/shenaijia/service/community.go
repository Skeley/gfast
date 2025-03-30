// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	api "github.com/tiger1103/gfast/v3/api/shenaijia/v1"
)

type (
	ICommunity interface {
		Search(ctx context.Context, req *api.CommunitySearchReq) (res *api.CommunityRes, err error)
		Add(ctx context.Context, req *api.CommunityAddReq) (res *api.CommunityAddRes, err error)
		Update(ctx context.Context, req *api.CommunityUpdateReq) (res *api.CommunityUpdateRes, err error)
	}
)

var (
	localCommunity ICommunity
)

func Community() ICommunity {
	if localCommunity == nil {
		panic("implement not found for interface ICommunity, forgot register?")
	}
	return localCommunity
}

func RegisterCommunity(i ICommunity) {
	localCommunity = i
}

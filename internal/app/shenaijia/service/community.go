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
		ListMajor(ctx context.Context, req *api.CommunityListMajorReq) (res *api.CommunityListMajorRes, err error)
		ListMinor(ctx context.Context, req *api.CommunityListMinorReq) (res *api.CommunityListMinorRes, err error)
		SysAdd(ctx context.Context, req *api.SysCommunityAddReq) (res *api.SysCommunityAddRes, err error)
		SysUpdate(ctx context.Context, req *api.SysCommunityUpdateReq) (res *api.SysCommunityUpdateRes, err error)
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

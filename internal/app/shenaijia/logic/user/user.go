package user

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "github.com/tiger1103/gfast/v3/api/shenaijia/v1"
	"github.com/tiger1103/gfast/v3/internal/app/shenaijia/dao"
	"github.com/tiger1103/gfast/v3/library/liberr"
)

func New() *sUser {
	return &sUser{}
}

type sUser struct{}

func (u *sUser) BindCommunity(ctx context.Context, req *v1.UserBindCommunityReq) (res *v1.UserBindCommunityRes, err error) {
	res = &v1.UserBindCommunityRes{}
	data := g.Map{
		dao.CommunityUnion.Columns().UserId:      req.UserId,
		dao.CommunityUnion.Columns().CommunityId: req.CommunityId,
	}
	err = g.Try(ctx, func(ctx context.Context) {
		_, e := dao.CommunityUnion.Ctx(ctx).Insert(data)
		liberr.ErrIsNil(ctx, e, "用户绑定小区失败")
	})
	return
}

func (u *sUser) UnbindCommunity(ctx context.Context, req *v1.UserUnbindCommunityReq) (res *v1.UserUnbindCommunityRes, err error) {
	res = &v1.UserUnbindCommunityRes{}

	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.CommunityUnion.Ctx(ctx)
		m.Where(dao.CommunityUnion.Columns().UserId, req.UserId)
		if len(req.CommunityId) > 0 {
			m.Where(dao.CommunityUnion.Columns().CommunityId, gconv.Uint64(req.CommunityId))
		}
		_, e := m.Delete()
		liberr.ErrIsNil(ctx, e, "用户解绑小区失败")
	})
	return
}

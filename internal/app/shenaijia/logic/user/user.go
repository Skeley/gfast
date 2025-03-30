package user

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "github.com/tiger1103/gfast/v3/api/shenaijia/v1"
	"github.com/tiger1103/gfast/v3/internal/app/shenaijia/dao"
	"github.com/tiger1103/gfast/v3/internal/app/shenaijia/model/entity"
	"github.com/tiger1103/gfast/v3/internal/app/shenaijia/service"
	"github.com/tiger1103/gfast/v3/library/liberr"
)

func init() {
	service.RegisterUser(New())
}

func New() *sUser {
	return &sUser{}
}

type sUser struct{}

func (s *sUser) BoundCommunity(ctx context.Context, req *v1.UserBoundCommunityReq) (res *v1.UserBoundCommunityRes, err error) {
	res = &v1.UserBoundCommunityRes{}
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.CommunityUnion.DB().Model(dao.CommunityUnion.Table(), "u").Safe().Ctx(ctx)
		m = m.InnerJoin(dao.Community.Table()+" c", "u.community_id=c.id").Where("u.user_id", req.UserId)
		var communityList []entity.Community
		err = m.Fields("c.*").Scan(&communityList)
		liberr.ErrIsNil(ctx, err, "获取小区列表失败")
		for _, community := range communityList {
			res.List = append(res.List, &v1.Community{
				Id:   community.Id,
				Pid:  community.Pid,
				Name: community.CommunityName,
			})
		}
		res.List, err = service.Community().FillParentName(ctx, res.List)
		liberr.ErrIsNil(ctx, err, "获取一级小区名失败")
	})
	return
}

func (s *sUser) BindCommunity(ctx context.Context, req *v1.UserBindCommunityReq) (res *v1.UserBindCommunityRes, err error) {
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

func (s *sUser) UnbindCommunity(ctx context.Context, req *v1.UserUnbindCommunityReq) (res *v1.UserUnbindCommunityRes, err error) {
	res = &v1.UserUnbindCommunityRes{}

	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.CommunityUnion.Ctx(ctx)
		m = m.Where(dao.CommunityUnion.Columns().UserId, req.UserId)
		if len(req.CommunityId) > 0 {
			m = m.Where(dao.CommunityUnion.Columns().CommunityId, gconv.Uint64(req.CommunityId))
		}
		_, e := m.Delete()
		liberr.ErrIsNil(ctx, e, "用户解绑小区失败")
	})
	return
}

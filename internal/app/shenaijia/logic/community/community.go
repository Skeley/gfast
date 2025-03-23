package community

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	api "github.com/tiger1103/gfast/v3/api/shenaijia/v1"
	"github.com/tiger1103/gfast/v3/internal/app/shenaijia/dao"
	"github.com/tiger1103/gfast/v3/internal/app/shenaijia/model/entity"
	"github.com/tiger1103/gfast/v3/library/liberr"
)

func New() *sCommunity {
	return &sCommunity{}
}

type sCommunity struct{}

func (s *sCommunity) ListMajor(ctx context.Context, req *api.CommunityListMajorReq) (res *api.CommunityListMajorRes, err error) {
	res = &api.CommunityListMajorRes{}
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.Community.Ctx(ctx).Where(dao.Community.Columns().MinorId, 0).Scan(&res.CommunityList)
		liberr.ErrIsNil(ctx, err, "获取一级小区列表失败")
	})
	return
}

func (s *sCommunity) ListMinor(ctx context.Context, req *api.CommunityListMinorReq) (res *api.CommunityListMinorRes, err error) {
	res = &api.CommunityListMinorRes{}
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.Community.Ctx(ctx).Where(dao.Community.Columns().MajorId, req.MajorId).Scan(&res.CommunityList)
		liberr.ErrIsNil(ctx, err, "获取二级小区列表失败")
	})
	return
}

func (s *sCommunity) SysAdd(ctx context.Context, req *api.SysCommunityAddReq) (res *api.SysCommunityAddRes, err error) {
	res = &api.SysCommunityAddRes{}
	data := entity.Community{
		MajorId:       req.MajorId,
		MinorId:       req.MinorId,
		CommunityName: req.Name,
	}
	err = g.Try(ctx, func(ctx context.Context) {
		_, err := dao.Community.Ctx(ctx).Insert(data)
		liberr.ErrIsNil(ctx, err, "添加小区失败")
	})
	return
}

func (s *sCommunity) SysUpdate(ctx context.Context, req *api.SysCommunityUpdateReq) (res *api.SysCommunityUpdateRes, err error) {
	res = &api.SysCommunityUpdateRes{}
	data := g.Map{
		dao.Community.Columns().CommunityName: req.Name,
	}
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			m := dao.Community.Ctx(ctx).TX(tx)
			m.Where(dao.Community.Columns().MajorId, req.MajorId)
			m.Where(dao.Community.Columns().MinorId, req.MinorId)
			_, e := m.Update(data)
			liberr.ErrIsNil(ctx, e, "修改小区名字失败")
		})
		return err
	})
	return
}

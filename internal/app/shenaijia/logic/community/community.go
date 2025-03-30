package community

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	api "github.com/tiger1103/gfast/v3/api/shenaijia/v1"
	"github.com/tiger1103/gfast/v3/internal/app/shenaijia/dao"
	"github.com/tiger1103/gfast/v3/library/liberr"
)

func New() *sCommunity {
	return &sCommunity{}
}

type sCommunity struct{}

func (s *sCommunity) Search(ctx context.Context, req *api.CommunitySearchReq) (res *api.CommunityRes, err error) {
	res = &api.CommunityRes{}
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.Community.Ctx(ctx)
		if len(req.Name) > 0 {
			m.Where("name LIKE ?", "%"+req.Name+"%")
		}
		if len(req.Parent) > 0 {
			m.Where("pid = ?", gconv.Uint64(req.Parent))
		}
		m.Scan(&res.CommunityList)
		liberr.ErrIsNil(ctx, err, "获取小区列表失败")
	})
	return
}

func (s *sCommunity) Add(ctx context.Context, req *api.CommunityAddReq) (res *api.CommunityAddRes, err error) {
	res = &api.CommunityAddRes{}
	data := g.Map{
		dao.Community.Columns().CommunityName: req.Name,
	}
	if len(req.Parent) > 0 {
		data[dao.Community.Columns().Pid] = gconv.Uint64(req.Parent)
	}
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			m := dao.Community.Ctx(ctx).TX(tx)
			_, e := m.Insert(data)
			liberr.ErrIsNil(ctx, e, "添加小区失败")
		})
		return err
	})
	return
}

func (s *sCommunity) Update(ctx context.Context, req *api.CommunityUpdateReq) (res *api.CommunityUpdateRes, err error) {
	res = &api.CommunityUpdateRes{}
	data := g.Map{
		dao.Community.Columns().CommunityName: req.Name,
	}
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			m := dao.Community.Ctx(ctx).TX(tx)
			m.Where(dao.Community.Columns().Id, req.Id)
			_, e := m.Update(data)
			liberr.ErrIsNil(ctx, e, "修改小区名字失败")
		})
		return err
	})
	return
}

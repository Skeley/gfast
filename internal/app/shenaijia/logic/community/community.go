package community

import (
	"context"
	api "github.com/tiger1103/gfast/v3/api/shenaijia/v1"
)

func New() *sCommunity {
	return &sCommunity{}
}

type sCommunity struct{}

func (s *sCommunity) ListAll(ctx context.Context, req *api.CommunityListAllReq) (res *api.CommunityRes, err error) {
	//res = &api.CommunityListMajorRes{}
	//err = g.Try(ctx, func(ctx context.Context) {
	//	err = dao.Community.Ctx(ctx).Where(dao.Community.Columns().MinorId, 0).Scan(&res.CommunityList)
	//	liberr.ErrIsNil(ctx, err, "获取一级小区列表失败")
	//})
	return
}

func (s *sCommunity) Search(ctx context.Context, req *api.CommunitySearchReq) (res *api.CommunityRes, err error) {
	//res = &api.CommunityListMinorRes{}
	//err = g.Try(ctx, func(ctx context.Context) {
	//	err = dao.Community.Ctx(ctx).Where(dao.Community.Columns().MajorId, req.MajorId).Scan(&res.CommunityList)
	//	liberr.ErrIsNil(ctx, err, "获取二级小区列表失败")
	//})
	return
}

func (s *sCommunity) List(ctx context.Context, req *api.CommunityListReq) (res *api.CommunityRes, err error) {
	//res = &api.SysCommunityAddRes{}
	//data := entity.Community{
	//	MajorId:       req.MajorId,
	//	MinorId:       req.MinorId,
	//	CommunityName: req.Name,
	//}
	//err = g.Try(ctx, func(ctx context.Context) {
	//	_, err := dao.Community.Ctx(ctx).Insert(data)
	//	liberr.ErrIsNil(ctx, err, "添加小区失败")
	//})
	return
}

func (s *sCommunity) Add(ctx context.Context, req *api.CommunityAddReq) (res *api.CommunityAddRes, err error) {
	//res = &api.SysCommunityUpdateRes{}
	//data := g.Map{
	//	dao.Community.Columns().CommunityName: req.Name,
	//}
	//err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
	//	err = g.Try(ctx, func(ctx context.Context) {
	//		m := dao.Community.Ctx(ctx).TX(tx)
	//		m.Where(dao.Community.Columns().MajorId, req.MajorId)
	//		m.Where(dao.Community.Columns().MinorId, req.MinorId)
	//		_, e := m.Update(data)
	//		liberr.ErrIsNil(ctx, e, "修改小区名字失败")
	//	})
	//	return err
	//})
	return
}

func (s *sCommunity) Update(ctx context.Context, req *api.CommunityUpdateReq) (res *api.CommunityUpdateRes, err error) {
	//res = &api.SysCommunityUpdateRes{}
	//data := g.Map{
	//	dao.Community.Columns().CommunityName: req.Name,
	//}
	//err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
	//	err = g.Try(ctx, func(ctx context.Context) {
	//		m := dao.Community.Ctx(ctx).TX(tx)
	//		m.Where(dao.Community.Columns().MajorId, req.MajorId)
	//		m.Where(dao.Community.Columns().MinorId, req.MinorId)
	//		_, e := m.Update(data)
	//		liberr.ErrIsNil(ctx, e, "修改小区名字失败")
	//	})
	//	return err
	//})
	return
}

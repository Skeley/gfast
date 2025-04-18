package community

import (
	"context"
	"errors"
	"github.com/go-sql-driver/mysql"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	api "github.com/tiger1103/gfast/v3/api/shenaijia/v1"
	"github.com/tiger1103/gfast/v3/internal/app/shenaijia/dao"
	"github.com/tiger1103/gfast/v3/internal/app/shenaijia/model/entity"
	"github.com/tiger1103/gfast/v3/internal/app/shenaijia/service"
	systemConsts "github.com/tiger1103/gfast/v3/internal/app/system/consts"
	"github.com/tiger1103/gfast/v3/library/liberr"
)

func init() {
	service.RegisterCommunity(New())
}

func New() *sCommunity {
	return &sCommunity{}
}

type sCommunity struct{}

func (s *sCommunity) getCommunityName(ctx context.Context, ids []uint) (res map[uint]string, err error) {
	res = make(map[uint]string)
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.Community.Ctx(ctx)
		var communityList []entity.Community
		m.WhereIn(dao.Community.Columns().Id, ids).Scan(&communityList)
		liberr.ErrIsNil(ctx, err, "获取小区列表失败")
		for _, community := range communityList {
			res[community.Id] = community.CommunityName
		}
	})
	return
}

func (s *sCommunity) FillParentName(ctx context.Context, list []*api.Community) ([]*api.Community, error) {
	var needParentNameIdList []uint
	for _, v := range list {
		if v.Pid != 0 {
			needParentNameIdList = append(needParentNameIdList, v.Pid)
		}
	}
	nameMap, e := s.getCommunityName(ctx, needParentNameIdList)
	if e != nil {
		return nil, e
	}
	for _, v := range list {
		if v.Pid != 0 {
			v.ParentName = nameMap[v.Pid]
		}
	}
	return list, nil
}

func (s *sCommunity) Search(ctx context.Context, req *api.CommunitySearchReq) (res *api.CommunityRes, err error) {
	res = &api.CommunityRes{}
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.Community.Ctx(ctx)
		if len(req.Name) > 0 {
			m = m.Where("community_name LIKE ?", "%"+req.Name+"%")
		}
		if len(req.Pid) > 0 {
			m = m.Where("pid = ?", gconv.Uint64(req.Pid))
		}
		res.Total, err = m.Count()

		if req.PageNum == 0 {
			req.PageNum = 1
		}
		res.CurrentPage = req.PageNum
		if req.PageSize == 0 {
			req.PageSize = systemConsts.PageSize
		}
		var communityList []entity.Community
		err = m.Page(req.PageNum, req.PageSize).Order(dao.Community.Columns().CommunityName + " asc").Scan(&communityList)
		liberr.ErrIsNil(ctx, err, "获取小区列表失败")
		for _, v := range communityList {
			res.List = append(res.List, &api.Community{
				Id:   v.Id,
				Pid:  v.Pid,
				Name: v.CommunityName,
			})
		}
		res.List, err = s.FillParentName(ctx, res.List)
		liberr.ErrIsNil(ctx, err, "获取一级小区名失败")
	})
	return
}

func (s *sCommunity) Add(ctx context.Context, req *api.CommunityAddReq) (res *api.CommunityAddRes, err error) {
	res = &api.CommunityAddRes{}
	data := g.Map{
		dao.Community.Columns().CommunityName: req.Name,
	}
	if len(req.Pid) > 0 {
		data[dao.Community.Columns().Pid] = gconv.Uint64(req.Pid)
	}
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			m := dao.Community.Ctx(ctx).TX(tx)
			_, e := m.Insert(data)
			var sqlErr *mysql.MySQLError
			if errors.As(e, &sqlErr) && sqlErr.Number == 1062 {
				g.Throw(errors.New("添加小区失败: 重复添加"))
			}
			// mysql.MySQLError{}
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
	if len(req.Pid) > 0 {
		data[dao.Community.Columns().Pid] = gconv.Uint64(req.Pid)
	}
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			m := dao.Community.Ctx(ctx).TX(tx)
			m = m.Where(dao.Community.Columns().Id, req.Id)
			_, e := m.Update(data)
			liberr.ErrIsNil(ctx, e, "修改小区失败")
		})
		return err
	})
	return
}

func (s *sCommunity) BoundUser(ctx context.Context, communityId uint) (userId uint64, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		var r *entity.CommunityUnion
		e := dao.CommunityUnion.Ctx(ctx).Where("community_id = ?", communityId).Scan(&r)
		liberr.ErrIsNil(ctx, e, "检索小区物业失败")
		if r == nil {
			g.Throw(errors.New("小区不存在"))
		}
		userId = uint64(r.UserId)
	})
	return
}

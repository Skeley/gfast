package project

import (
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"slices"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	api "github.com/tiger1103/gfast/v3/api/shenaijia/v1"
	"github.com/tiger1103/gfast/v3/internal/app/shenaijia/dao"
	"github.com/tiger1103/gfast/v3/internal/app/shenaijia/model/entity"
	"github.com/tiger1103/gfast/v3/internal/app/shenaijia/service"
	systemConsts "github.com/tiger1103/gfast/v3/internal/app/system/consts"
	sysModel "github.com/tiger1103/gfast/v3/internal/app/system/model"
	"github.com/tiger1103/gfast/v3/library/liberr"
)

func init() {
	service.RegisterProject(New())
}

func New() *sProject {
	return &sProject{}
}

type sProject struct{}

func (s *sProject) List(ctx context.Context, user *sysModel.LoginUserRes, req *api.ProjectListReq) (res *api.ProjectListRes, err error) {
	res = &api.ProjectListRes{}
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.Project.Ctx(ctx).InnerJoin(dao.Community.Table(), "c",
			fmt.Sprintf("%s.%s=c.%s",
				dao.Project.Table(),
				dao.Project.Columns().CommunityId,
				dao.Community.Columns().Id))
		if slices.Index(user.UserTypes, req.UserType) == -1 {
			g.Throw(errors.New("非法请求, 用户身份无效"))
		}
		m = m.Where(dao.Project.Columns().Valid, true)
		switch req.UserType {
		case 1:
			m = m.Where(dao.Project.Columns().Manager, user.Id)
		case 2:
			m = m.Where(dao.Project.Columns().Creator, user.Id)
		case 3:
			m = m.Where(dao.Project.Columns().Associate, user.Id)
		}
		res.Total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "获取项目列表失败")

		if req.PageNum == 0 {
			req.PageNum = 1
		}
		res.CurrentPage = req.PageNum
		if req.PageSize == 0 {
			req.PageSize = systemConsts.PageSize
		}
		err = m.Fields("project.*").Fields(dao.Community.Columns().CommunityName).
			Page(req.PageNum, req.PageSize).Order(dao.Project.Columns().StartDate + " desc").
			Scan(&res.List)
		liberr.ErrIsNil(ctx, err, "获取项目列表失败")
	})
	return
}

func (s *sProject) SysList(ctx context.Context, req *api.SysProjectSearchReq) (res *api.SysProjectSearchRes, err error) {
	res = &api.SysProjectSearchRes{}
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.Project.Ctx(ctx).InnerJoin(dao.Community.Table(), "c", "project.community_id = c.id")
		m = m.Where(dao.Project.Columns().Valid, true)
		if len(req.CommunityId) > 0 {
			m = m.Where(dao.Project.Columns().CommunityId, gconv.Uint(req.CommunityId))
		}
		if len(req.Manager) > 0 {
			m = m.Where(entity.Project{Manager: gconv.Uint(req.Manager)})
		}
		if len(req.Associate) > 0 {
			m = m.Where(entity.Project{Associate: gconv.Uint(req.Associate)})
		}
		if len(req.CommunityId) > 0 {
			m = m.Where(entity.Project{CommunityId: gconv.Uint(req.CommunityId)})
		}
		if len(req.StartDateRange) > 0 {
			m = m.WhereBetween(dao.Project.Columns().StartDate, req.StartDateRange[0], req.StartDateRange[1])
		}
		if len(req.CompletionDateRange) > 0 {
			m = m.WhereBetween(dao.Project.Columns().CompletionDate, req.CompletionDateRange[0], req.CompletionDateRange[1])
		}
		res.Total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "获取项目列表失败")

		if req.PageNum == 0 {
			req.PageNum = 1
		}
		res.CurrentPage = req.PageNum
		if req.PageSize == 0 {
			req.PageSize = systemConsts.PageSize
		}
		err = m.Fields("project.*").
			Fields(
				"c.id as community_id",
				"c.pid as community_pid",
				"c.community_name").
			Page(req.PageNum, req.PageSize).Order(dao.Project.Columns().StartDate + " desc").
			Scan(&res.List)
		liberr.ErrIsNil(ctx, err, "获取项目列表失败")
	})
	return
}

func (s *sProject) SysAdd(ctx context.Context, req *api.SysProjectAddReq) (res *api.SysProjectAddRes, err error) {
	res = &api.SysProjectAddRes{}
	err = g.Try(ctx, func(ctx context.Context) {
		data := entity.Project{
			ProjectName: req.ProjectName,
			CommunityId: req.CommunityId,
			Progress:    req.Progress,
			Manager:     req.Manager,
		}
		if len(req.StartDate) > 0 {
			data.StartDate = gtime.New(req.StartDate)
		}
		if len(req.EstimatedCompletionDate) > 0 {
			data.EstimatedCompletionDate = gtime.New(req.EstimatedCompletionDate)
		}
		if len(req.InspectionReport) > 0 {
			data.InspectionReport = req.InspectionReport
		}
		_, err := dao.Project.Ctx(ctx).Insert(data)
		liberr.ErrIsNil(ctx, err, "添加项目失败")
	})
	return
}

func (s *sProject) SysEdit(ctx context.Context, req *api.SysProjectEditReq) (res *api.SysProjectEditRes, err error) {
	res = &api.SysProjectEditRes{}
	data := g.Map{
		dao.Project.Columns().ProjectName:      req.ProjectName,
		dao.Project.Columns().CommunityId:      req.CommunityId,
		dao.Project.Columns().Progress:         req.Progress,
		dao.Project.Columns().InspectionReport: req.InspectionReport,
		dao.Project.Columns().AcceptanceReport: req.AcceptanceReport,
		dao.Project.Columns().Manager:          req.Manager,
		dao.Project.Columns().Associate:        req.Associate,
	}
	if len(req.StartDate) > 0 {
		data[dao.Project.Columns().StartDate] = gtime.New(req.StartDate)
	}
	if len(req.EstimatedCompletionDate) > 0 {
		data[dao.Project.Columns().EstimatedCompletionDate] = gtime.New(req.EstimatedCompletionDate)
	}
	if len(req.CompletionDate) > 0 {
		data[dao.Project.Columns().CompletionDate] = gtime.New(req.CompletionDate)
	}
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			_, e := dao.Project.Ctx(ctx).TX(tx).WherePri(req.ProjectId).Update(data)
			liberr.ErrIsNil(ctx, e, "修改项目失败")
		})
		return err
	})
	return
}

func (s *sProject) Delete(ctx context.Context, req *api.ProjectDeleteReq) (res *api.ProjectDeleteRes, err error) {
	res = &api.ProjectDeleteRes{}
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			_, e := dao.Project.Ctx(ctx).TX(tx).WherePri(req.ProjectId).Update(
				g.Map{dao.Project.Columns().Valid: false})
			liberr.ErrIsNil(ctx, e, "软删除项目失败")
		})
		return err
	})
	return
}

func (s *sProject) Audit(ctx context.Context, req *api.SysProjectAuditReq) (res *api.SysProjectAuditRes, err error) {
	res = &api.SysProjectAuditRes{}
	err = g.Try(ctx, func(ctx context.Context) {
		rec, e := dao.Project.Ctx(ctx).WherePri(req.Id).Fields(dao.Project.Columns().Audited).One()
		liberr.ErrIsNil(ctx, e, "获取项目信息失败")

		if rec["audited"].Bool() {
			g.Throw(errors.New("项目已审核，无需重复审核"))
		}

		data := g.Map{
			dao.Project.Columns().Audited: true,
		}
		if len(req.Associate) > 0 {
			data[dao.Project.Columns().Associate] = gconv.Uint(req.Associate)
		}
		_, e = dao.Project.Ctx(ctx).WherePri(req.Id).Update(data)
		liberr.ErrIsNil(ctx, e, "审核失败")
	})
	return
}

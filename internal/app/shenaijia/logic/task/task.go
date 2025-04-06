package task

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "github.com/tiger1103/gfast/v3/api/shenaijia/v1"
	"github.com/tiger1103/gfast/v3/internal/app/shenaijia/dao"
	"github.com/tiger1103/gfast/v3/internal/app/shenaijia/model/entity"
	"github.com/tiger1103/gfast/v3/internal/app/shenaijia/service"
	systemConsts "github.com/tiger1103/gfast/v3/internal/app/system/consts"
	"github.com/tiger1103/gfast/v3/library/liberr"
)

func init() {
	service.RegisterTask(New())
}

func New() *sTask {
	return &sTask{}
}

type sTask struct{}

func (s *sTask) SearchType(ctx context.Context, req *v1.TempletTypeSearchReq) (res *v1.TempletTypeSearchRes, err error) {
	res = &v1.TempletTypeSearchRes{}
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.TaskType.Ctx(ctx)
		if len(req.Name) > 0 {
			m = m.Where("name LIKE ?", "%"+req.Name+"%")
		}
		if req.Standard != nil {
			m = m.Where(dao.TaskType.Columns().Standard, req.Standard)
		}
		res.Total, err = m.Count()

		if req.PageNum == 0 {
			req.PageNum = 1
		}
		res.CurrentPage = req.PageNum
		if req.PageSize == 0 {
			req.PageSize = systemConsts.PageSize
		}
		err = m.Page(req.PageNum, req.PageSize).Scan(&res.List)
		liberr.ErrIsNil(ctx, err, "获取类型列表失败")
	})
	return
}

func (s *sTask) AddType(ctx context.Context, req *v1.TempletTypeAddReq) (res *v1.TempletTypeAddRes, err error) {
	res = &v1.TempletTypeAddRes{}
	data := g.Map{
		dao.TaskType.Columns().Name:     req.Name,
		dao.TaskType.Columns().Standard: req.Standard,
	}
	err = g.Try(ctx, func(ctx context.Context) {
		_, e := dao.TaskType.Ctx(ctx).Insert(data)
		liberr.ErrIsNil(ctx, e, "新增类型失败")
	})
	return
}

func (s *sTask) UpdateType(ctx context.Context, req *v1.TempletTypeUpdateReq) (res *v1.TempletTypeUpdateRes, err error) {
	res = &v1.TempletTypeUpdateRes{}
	data := g.Map{
		dao.TaskType.Columns().Name:     req.Name,
		dao.TaskType.Columns().Standard: req.Standard,
	}
	err = g.Try(ctx, func(ctx context.Context) {
		_, e := dao.TaskType.Ctx(ctx).WherePri(req.Id).Update(data)
		liberr.ErrIsNil(ctx, e, "修改类型失败")
	})
	return
}

func (s *sTask) DeleteType(ctx context.Context, req *v1.TempletTypeDeleteReq) (res *v1.TempletTypeDeleteRes, err error) {
	res = &v1.TempletTypeDeleteRes{}
	err = g.Try(ctx, func(ctx context.Context) {
		// check have templet
		var t *entity.TaskType
		dao.TaskType.Ctx(ctx).WherePri(req.Id).Scan(&t)
		if t == nil {
			return
		}

		hasTemplet, e := s.HasTemplet(ctx, t.Id)
		liberr.ErrIsNil(ctx, e, "删除类型失败")
		if hasTemplet {
			g.Throw(errors.New("被模板依赖, 无法删除"))
		}
		_, e = dao.TaskType.Ctx(ctx).WherePri(req.Id).Delete()
		liberr.ErrIsNil(ctx, e, "删除类型失败")
	})
	return
}

func (s *sTask) HasTemplet(ctx context.Context, templetType uint) (res bool, err error) {
	res = false
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.TaskTemplet.DB().Model(dao.TaskTemplet.Table(), "p").Safe().Ctx(ctx)
		m = m.InnerJoin(dao.TaskType.Table()+" t", "t.id = p.type")
		m = m.Where("p.type = ?", templetType)
		total, e := m.Count()
		liberr.ErrIsNil(ctx, e, "获取任务模板失败")
		res = total > 0
	})
	return
}

func (s *sTask) SearchTemplet(ctx context.Context, req *v1.TaskTempletSearchReq) (res *v1.TaskTempletSearchRes, err error) {
	res = &v1.TaskTempletSearchRes{}
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.TaskTemplet.DB().Model(dao.TaskTemplet.Table(), "p").Safe().Ctx(ctx)
		m = m.InnerJoin(dao.TaskType.Table()+" t", "t.id = p.type")
		if len(req.Name) > 0 {
			m = m.Where("p.name LIKE ?", "%"+req.Name+"%")
		}
		if len(req.Type) > 0 {
			m = m.Where("p.type = ?", gconv.Uint(req.Type))
		}
		res.Total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "获取任务模板失败")

		if req.PageNum == 0 {
			req.PageNum = 1
		}
		res.CurrentPage = req.PageNum
		if req.PageSize == 0 {
			req.PageSize = systemConsts.PageSize
		}

		err = m.Fields("p.*", "t.standard").Page(req.PageNum, req.PageSize).Scan(&res.List)
		liberr.ErrIsNil(ctx, err, "获取任务模板失败")
	})
	return
}

func (s *sTask) AddTemplet(ctx context.Context, req *v1.TaskTempletAddReq) (res *v1.TaskTempletAddRes, err error) {
	res = &v1.TaskTempletAddRes{}
	data := g.Map{
		dao.TaskTemplet.Columns().Name: req.Name,
		dao.TaskTemplet.Columns().Type: req.Type,
	}
	err = g.Try(ctx, func(ctx context.Context) {
		_, e := dao.TaskTemplet.Ctx(ctx).Insert(data)
		liberr.ErrIsNil(ctx, e, "添加模板失败")
	})
	return
}

func (s *sTask) UpdateTemplet(ctx context.Context, req *v1.TaskTempletUpdateReq) (res *v1.TaskTempletUpdateRes, err error) {
	res = &v1.TaskTempletUpdateRes{}
	data := g.Map{
		dao.TaskTemplet.Columns().Name: req.Name,
		dao.TaskTemplet.Columns().Type: req.Type,
	}
	err = g.Try(ctx, func(ctx context.Context) {
		_, e := dao.TaskTemplet.Ctx(ctx).WherePri(req.Id).Update(data)
		liberr.ErrIsNil(ctx, e, "修改模板失败")
	})
	return
}

func (s *sTask) DeleteTemplet(ctx context.Context, req *v1.TaskTempletDeleteReq) (res *v1.TaskTempletDeleteRes, err error) {
	res = &v1.TaskTempletDeleteRes{}
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			// 删除模板相关 stage
			//_, e := dao.TaskStage.Ctx(ctx).TX(tx).Where(dao.TaskStage.Columns().TempletId, req.Id).Delete()
			//liberr.ErrIsNil(ctx, e, "删除模板失败")
			_, e := dao.TaskTemplet.Ctx(ctx).WherePri(req.Id).Delete()
			liberr.ErrIsNil(ctx, e, "删除模板失败")
		})
		return err
	})
	return
}

func (s *sTask) TempletSetFlow(ctx context.Context, req *v1.TaskTempletSetFlowReq) (res *v1.TaskTempletSetFlowRes, err error) {
	res = &v1.TaskTempletSetFlowRes{}
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			stageIds, e := s.listStageIds(ctx, req.TempletId)
			liberr.ErrIsNil(ctx, e, "删除旧流程失败, 流程终止")
			for _, stageId := range stageIds {
				e = s.deleteSteps(ctx, tx, stageId)
				liberr.ErrIsNil(ctx, e, "删除旧流程失败, 流程终止")
			}
			e = s.deleteStages(ctx, tx, req.TempletId)
			liberr.ErrIsNil(ctx, e, "删除旧流程失败, 流程终止")
			// 设置新流程
			stageIds, e = s.addStages(ctx, tx, req.TempletId, req.Flow.Stages)
			liberr.ErrIsNil(ctx, e, "添加流程失败: 创建stage失败")
			for i, stage := range req.Flow.Stages {
				e = s.addSteps(ctx, tx, stageIds[i], stage.Steps)
				liberr.ErrIsNil(ctx, e, "添加流程失败: 创建step失败")
			}
		})
		return err
	})
	return
}

func (s *sTask) TempletGetFlow(ctx context.Context, req *v1.TaskTempletGetFlowReq) (res *v1.TaskTempletGetFlowRes, err error) {
	res = &v1.TaskTempletGetFlowRes{
		Flow: &v1.TaskFlow{},
	}
	err = g.Try(ctx, func(ctx context.Context) {
		stages, e := s.listStage(ctx, req.TempletId)
		liberr.ErrIsNil(ctx, e, "获取流程失败")
		for _, stage := range stages {
			stageApi := v1.TaskStage{
				Name: stage.Name,
				Icon: stage.Icon,
			}
			steps, e := s.listStep(ctx, stage.Id)
			liberr.ErrIsNil(ctx, e, "获取流程失败")
			for _, step := range steps {
				stageApi.Steps = append(stageApi.Steps, &v1.TaskStep{
					Name:    step.Name,
					Comment: step.Comment,
				})
			}
			res.Flow.Stages = append(res.Flow.Stages, &stageApi)
		}
	})
	if err != nil {
		return nil, err
	}
	return
}

func (s *sTask) ListStage(ctx context.Context, req *v1.TaskListStageReq) (res *v1.TaskListStageRes, err error) {
	res = &v1.TaskListStageRes{}
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.TaskStage.Ctx(ctx).
			Where(dao.TaskStage.Columns().TempletId, req.TempletId).
			Where(dao.TaskStage.Columns().TaskId, -1).
			OrderAsc(dao.TaskStage.Columns().Position).Scan(&res.List)
		liberr.ErrIsNil(ctx, err, "获取任务阶段失败")
	})
	return
}

func (s *sTask) listStageIds(ctx context.Context, templetId uint) (stageIds []uint, err error) {
	stages, e := s.listStage(ctx, templetId)
	if e != nil {
		return nil, e
	}
	for _, stage := range stages {
		stageIds = append(stageIds, stage.Id)
	}
	return
}

func (s *sTask) listStage(ctx context.Context, templetId uint) (stages []entity.TaskStage, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		e := dao.TaskStage.Ctx(ctx).
			Where(dao.TaskStage.Columns().TempletId, templetId).
			Where(dao.TaskStage.Columns().TaskId, -1).
			OrderAsc(dao.TaskStage.Columns().Position).Scan(&stages)
		liberr.ErrIsNil(ctx, e, "获取任务阶段失败")
	})
	return
}

func (s *sTask) addStages(ctx context.Context, tx gdb.TX, templetId uint, stages []*v1.TaskStage) (ids []uint, err error) {
	dataList := g.List{}
	for i, v := range stages {
		dataList = append(dataList, g.Map{
			dao.TaskStage.Columns().TempletId: templetId,
			dao.TaskStage.Columns().Name:      v.Name,
			dao.TaskStage.Columns().Icon:      v.Icon,
			dao.TaskStage.Columns().Position:  i,
		})
	}
	err = g.Try(ctx, func(ctx context.Context) {
		insertRes, e := dao.TaskStage.Ctx(ctx).TX(tx).Data(dataList).Insert()
		liberr.ErrIsNil(ctx, e, "添加阶段失败")
		firstId, _ := insertRes.LastInsertId()
		rowCnt, _ := insertRes.RowsAffected()
		for i := int64(0); i < rowCnt; i++ {
			ids = append(ids, uint(firstId+i))
		}
	})
	return
}

func (s *sTask) deleteStages(ctx context.Context, tx gdb.TX, templetId uint) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, e := dao.TaskStage.Ctx(ctx).TX(tx).
			Where(dao.TaskStage.Columns().TempletId, templetId).
			Where(dao.TaskStage.Columns().TaskId, -1).
			Delete()
		liberr.ErrIsNil(ctx, e, "删除模板stage失败")
	})
	return nil
}

func (s *sTask) AddStage(ctx context.Context, req *v1.TaskAddStageReq) (res *v1.TaskAddStageRes, err error) {
	res = &v1.TaskAddStageRes{}
	data := g.Map{
		dao.TaskStage.Columns().TempletId: req.TempletId,
		dao.TaskStage.Columns().Name:      req.Name,
		dao.TaskStage.Columns().Icon:      req.Icon,
		dao.TaskStage.Columns().Position:  req.Position,
	}
	err = g.Try(ctx, func(ctx context.Context) {
		_, e := dao.TaskStage.Ctx(ctx).Insert(data)
		liberr.ErrIsNil(ctx, e, "添加阶段失败")
	})
	return
}

func (s *sTask) UpdateStage(ctx context.Context, req *v1.TaskUpdateStageReq) (res *v1.TaskUpdateStageRes, err error) {
	res = &v1.TaskUpdateStageRes{}
	data := g.Map{
		dao.TaskStage.Columns().Name: req.Name,
		dao.TaskStage.Columns().Icon: req.Icon,
	}
	err = g.Try(ctx, func(ctx context.Context) {
		_, e := dao.TaskStage.Ctx(ctx).WherePri(req.StageId).Update(data)
		liberr.ErrIsNil(ctx, e, "更新阶段失败")
	})
	return
}

func (s *sTask) DeleteStage(ctx context.Context, req *v1.TaskDeleteStageReq) (res *v1.TaskDeleteStageRes, err error) {
	res = &v1.TaskDeleteStageRes{}
	err = g.Try(ctx, func(ctx context.Context) {
		_, e := dao.TaskStage.Ctx(ctx).WherePri(req.StageId).Delete()
		liberr.ErrIsNil(ctx, e, "删除阶段失败")
	})
	return
}

func (s *sTask) ListStep(ctx context.Context, req *v1.TaskListStepReq) (res *v1.TaskListStepRes, err error) {
	res = &v1.TaskListStepRes{}
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.TaskStep.Ctx(ctx).
			Where(dao.TaskStep.Columns().StageId, req.StageId).
			Where(dao.TaskStep.Columns().TaskId, req.TaskId).
			Order(dao.TaskStep.Columns().Position).Scan(&res.List)
		liberr.ErrIsNil(ctx, err, "获取任务步骤失败")
	})
	return
}

func (s *sTask) listStep(ctx context.Context, stageId uint) (res []*entity.TaskStep, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.TaskStep.Ctx(ctx).
			Where(dao.TaskStep.Columns().TaskId, -1).
			Where(dao.TaskStep.Columns().StageId, stageId).
			Scan(&res)
		liberr.ErrIsNil(ctx, err, "获取步骤失败")
	})
	return
}

func (s *sTask) addSteps(ctx context.Context, tx gdb.TX, stageId uint, steps []*v1.TaskStep) (err error) {
	dataList := g.List{}
	for i, v := range steps {
		dataList = append(dataList, g.Map{
			dao.TaskStep.Columns().StageId:  stageId,
			dao.TaskStep.Columns().Name:     v.Name,
			dao.TaskStep.Columns().Comment:  v.Comment,
			dao.TaskStep.Columns().Position: i,
		})
	}
	err = g.Try(ctx, func(ctx context.Context) {
		_, e := dao.TaskStep.Ctx(ctx).TX(tx).Data(dataList).Insert()
		liberr.ErrIsNil(ctx, e, "添加步骤失败")
	})
	return nil
}

func (s *sTask) deleteSteps(ctx context.Context, tx gdb.TX, stageId uint) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, e := dao.TaskStep.Ctx(ctx).TX(tx).
			Where(dao.TaskStep.Columns().StageId, stageId).
			Where(dao.TaskStep.Columns().TaskId, -1).
			Delete()
		liberr.ErrIsNil(ctx, e, "删除模板step失败")
	})
	return nil
}

func (s *sTask) AddStep(ctx context.Context, req *v1.TaskAddStepReq) (res *v1.TaskAddStepRes, err error) {
	res = &v1.TaskAddStepRes{}
	data := g.Map{
		dao.TaskStep.Columns().StageId: req.StageId,

		dao.TaskStep.Columns().Name:     req.StepName,
		dao.TaskStep.Columns().Position: 0,
	}
	if len(req.Comment) > 0 {
		data[dao.TaskStep.Columns().Comment] = req.Comment
	}
	if len(req.EstimatedCompletionDate) > 0 {
		data[dao.TaskStep.Columns().EstimatedCompletionDate] = req.EstimatedCompletionDate
	}
	err = g.Try(ctx, func(ctx context.Context) {
		_, e := dao.TaskStep.Ctx(ctx).Insert(data)
		liberr.ErrIsNil(ctx, e, "添加步骤失败")
	})
	return
}

func (s *sTask) UpdateStep(ctx context.Context, req *v1.TaskUpdateStepReq) (res *v1.TaskAddStepRes, err error) {
	res = &v1.TaskAddStepRes{}
	data := g.Map{
		dao.TaskStep.Columns().Name: req.StepName,
	}
	if len(req.Comment) > 0 {
		data[dao.TaskStep.Columns().Comment] = req.Comment
	}
	if len(req.EstimatedCompletionDate) > 0 {
		data[dao.TaskStep.Columns().EstimatedCompletionDate] = req.EstimatedCompletionDate
	}
	if req.State != nil {
		data[dao.TaskStep.Columns().State] = req.State
	}
	err = g.Try(ctx, func(ctx context.Context) {
		_, e := dao.TaskStep.Ctx(ctx).WherePri(req.StepId).Update(data)
		liberr.ErrIsNil(ctx, e, "添加步骤失败")
	})
	return
}

func (s *sTask) DeleteStep(ctx context.Context, req *v1.TaskDeleteStepReq) (res *v1.TaskDeleteStepRes, err error) {
	res = &v1.TaskDeleteStepRes{}
	err = g.Try(ctx, func(ctx context.Context) {
		_, e := dao.TaskStep.Ctx(ctx).WherePri(req.StepId).Delete()
		liberr.ErrIsNil(ctx, e, "删除步骤失败")
	})
	return
}

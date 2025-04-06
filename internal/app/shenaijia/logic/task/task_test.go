package task

import (
	"context"
	"encoding/json"
	"fmt"
	v1 "github.com/tiger1103/gfast/v3/api/shenaijia/v1"
	"testing"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
)

func Test_sTask_TempletSetFlow(t *testing.T) {
	req := &v1.TaskTempletSetFlowReq{}
	req.TempletId = 3
	req.Flow = v1.TaskFlow{Stages: []*v1.TaskStage{
		&v1.TaskStage{
			Name: "S1",
			Icon: "https://sss.com",
			Steps: []*v1.TaskStep{
				&v1.TaskStep{
					Name:    "S1T1",
					Comment: "comment",
				},
				&v1.TaskStep{
					Name:    "S1T2",
					Comment: "comment",
				},
				&v1.TaskStep{
					Name:    "S1T3",
					Comment: "comment",
				},
				&v1.TaskStep{
					Name:    "S1T4",
					Comment: "comment",
				},
			},
		},
		&v1.TaskStage{
			Name: "S2",
			Icon: "https://sss.com",
			Steps: []*v1.TaskStep{
				&v1.TaskStep{
					Name:    "S2T1",
					Comment: "comment",
				},
				&v1.TaskStep{
					Name:    "S2T2",
					Comment: "comment",
				},
			},
		},
		&v1.TaskStage{
			Name: "S3",
			Icon: "https://sss.com",
			Steps: []*v1.TaskStep{
				&v1.TaskStep{
					Name:    "S3T1",
					Comment: "comment",
				},
			},
		},
	}}
	ctx := context.Background()

	s := &sTask{}
	_, err := s.TempletSetFlow(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_sTask_TempletGetFlow(t *testing.T) {
	req := &v1.TaskTempletGetFlowReq{}
	req.TempletId = 3
	ctx := context.Background()

	s := &sTask{}
	res, err := s.TempletGetFlow(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	jsonData, err := json.Marshal(res)
	if err != nil {
		fmt.Println("JSON 序列化出错:", err)
		return
	}
	fmt.Println(string(jsonData))
}

func Test_sTask_ListStage(t *testing.T) {
	req := &v1.TaskListStageReq{}
	req.TempletId = 2

	ctx := context.Background()
	s := &sTask{}

	res, err := s.ListStage(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	for _, t := range res.List {
		fmt.Println(t)
	}
}

func Test_sTask_DeleteStage(t *testing.T) {
	req := &v1.TaskDeleteStageReq{}
	req.StageId = 1

	ctx := context.Background()
	s := &sTask{}

	_, err := s.DeleteStage(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_sTask_UpdateStage(t *testing.T) {
	req := &v1.TaskUpdateStageReq{}
	req.StageId = 1
	req.Name = "材料准备"
	req.Icon = "https://www.tt.com"

	ctx := context.Background()
	s := &sTask{}

	_, err := s.UpdateStage(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_sTask_AddStage(t *testing.T) {
	req := &v1.TaskAddStageReq{}
	req.Name = "施工"
	req.TempletId = 2
	req.Icon = "https://www.tt.com"
	req.Position = 2

	ctx := context.Background()
	s := &sTask{}

	_, err := s.AddStage(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_sTask_SearchTemplet(t *testing.T) {
	req := &v1.TaskTempletSearchReq{}
	req.PageSize = 10
	req.Name = "外墙"
	req.Type = "3"

	ctx := context.Background()
	s := &sTask{}

	res, err := s.SearchTemplet(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	for _, t := range res.List {
		fmt.Println(t)
	}
}

func Test_sTask_DeleteTemplet(t *testing.T) {
	req := &v1.TaskTempletDeleteReq{}
	req.Id = 1
	ctx := context.Background()
	s := &sTask{}

	_, err := s.DeleteTemplet(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_sTask_UpdateTemplet(t *testing.T) {
	req := &v1.TaskTempletUpdateReq{}
	req.Id = 2
	req.Type = 3
	req.Name = "外墙防水"
	ctx := context.Background()
	s := &sTask{}

	_, err := s.UpdateTemplet(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_sTask_AddTemplet(t *testing.T) {
	req := &v1.TaskTempletAddReq{}
	req.Type = 3
	req.Name = "防水"
	ctx := context.Background()
	s := &sTask{}

	_, err := s.AddTemplet(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_sTask_SearchType(t *testing.T) {
	req := &v1.TempletTypeSearchReq{}
	req.PageSize = 10
	req.Name = "水"
	ctx := context.Background()

	s := &sTask{}
	res, err := s.SearchType(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	for _, taskType := range res.List {
		fmt.Println(taskType)
	}
}

func Test_sTask_AddType(t *testing.T) {
	req := &v1.TempletTypeAddReq{}
	req.Standard = true
	req.Name = "防水"
	ctx := context.Background()

	s := &sTask{}
	_, err := s.AddType(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_sTask_UpdateType(t *testing.T) {
	req := &v1.TempletTypeUpdateReq{}
	req.Id = 1
	req.Standard = true
	req.Name = "防水"
	ctx := context.Background()

	s := &sTask{}
	_, err := s.UpdateType(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_sTask_DeleteType(t *testing.T) {
	req := &v1.TempletTypeDeleteReq{}
	req.Id = 3
	ctx := context.Background()

	s := &sTask{}
	_, err := s.DeleteType(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
}

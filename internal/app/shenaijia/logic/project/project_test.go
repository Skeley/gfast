package project

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	api "github.com/tiger1103/gfast/v3/api/shenaijia/v1"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
)

func Test_sSysProject_List(t *testing.T) {
	ctx := context.Background()
	req := api.SysProjectSearchReq{}

	s := &sProject{}
	gotRes, _ := s.SysList(ctx, &req)
	for _, project := range gotRes.List {
		s, _ := json.Marshal(project)
		fmt.Println(string(s))
	}
}

func Test_sSysProject_Audit(t *testing.T) {
	ctx := context.Background()
	req := api.SysProjectAuditReq{}
	req.Id = 2
	req.Associate = "1"

	s := &sProject{}
	_, err := s.Audit(ctx, &req)
	if err != nil {
		t.Error(err)
	}
}

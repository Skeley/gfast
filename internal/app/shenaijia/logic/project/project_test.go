package project

import (
	"context"
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
		fmt.Println(project)
	}
}

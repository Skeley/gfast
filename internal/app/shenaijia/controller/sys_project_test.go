package controller

import (
	"context"
	"encoding/json"
	"fmt"
	api "github.com/tiger1103/gfast/v3/api/shenaijia/v1"
	"testing"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/tiger1103/gfast/v3/internal/app/boot"
)

func Test_sysProjectController_ProjectList(t *testing.T) {
	ctx := context.Background()
	req := api.SysProjectSearchReq{}

	gotRes, _ := SysProject.ProjectList(ctx, &req)
	for _, project := range gotRes.List {
		s, _ := json.Marshal(project)
		fmt.Println(string(s))
	}
}

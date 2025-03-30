package user

import (
	"context"
	"fmt"
	"testing"

	api "github.com/tiger1103/gfast/v3/api/shenaijia/v1"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/tiger1103/gfast/v3/internal/app/shenaijia/logic/community"
)

func Test_sUser_BoundCommunity(t *testing.T) {
	ctx := context.Background()
	req := api.UserBoundCommunityReq{
		UserId: 1,
	}
	s := &sUser{}
	gotRes, err := s.BoundCommunity(ctx, &req)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, community := range gotRes.List {
		fmt.Println(community)
	}
}

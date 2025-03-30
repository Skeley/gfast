package community

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/gogf/gf/v2/frame/g"
	api "github.com/tiger1103/gfast/v3/api/shenaijia/v1"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
)

func Test_sCommunity_Search(t *testing.T) {
	ctx := context.Background()
	req := api.CommunitySearchReq{
		Meta: g.Meta{},
		Pid:  "",
		Name: "",
	}
	req.PageNum = 1
	req.PageSize = 100

	s := &sCommunity{}
	gotRes, err := s.Search(ctx, &req)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, community := range gotRes.List {
		fmt.Println(community)
	}
}

func Test_sCommunity_Add(t *testing.T) {
	t.Run("add", func(t *testing.T) {
		ctx := context.Background()
		req := api.CommunityAddReq{
			Pid:  "12",
			Name: "fsk2",
		}

		s := &sCommunity{}
		_, err := s.Add(ctx, &req)
		assert.NoError(t, err)
	})
}

func Test_sCommunity_Update(t *testing.T) {
	t.Run("update", func(t *testing.T) {
		ctx := context.Background()
		req := api.CommunityUpdateReq{
			Id:   13,
			Pid:  "0",
			Name: "fsk2",
		}

		s := &sCommunity{}
		_, err := s.Update(ctx, &req)
		assert.NoError(t, err)
	})
}

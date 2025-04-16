package sysUser

import (
	"context"
	"testing"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
)

func Test_sSysUser_GetUserByMobile(t *testing.T) {
	ctx := context.Background()
	s := &sSysUser{
		casBinUserPrefix: "u_",
	}
	user, err := s.GetUserByMobile(ctx, "123")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(user)
}

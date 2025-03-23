package user

import (
	"context"
	"fmt"
	"testing"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
)

func Test_sUser_GetUserType(t *testing.T) {
	ctx := context.Background()

	s := &sUser{}
	userTypes, _ := s.GetUserType(ctx, 31)
	fmt.Println(userTypes)
}

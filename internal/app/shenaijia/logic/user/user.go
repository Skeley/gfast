package user

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/tiger1103/gfast/v3/internal/app/shenaijia/dao"
	"github.com/tiger1103/gfast/v3/internal/app/shenaijia/model/entity"
	"github.com/tiger1103/gfast/v3/library/liberr"
	"strings"
)

func New() *sUser {
	return &sUser{}
}

type sUser struct{}

func (u *sUser) GetUserType(ctx context.Context, userId uint64) (userTypes []string, err error) {
	userInfo := &entity.UserInfo{}
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.UserInfo.Ctx(ctx).Where(dao.UserInfo.Columns().Id, userId).Scan(userInfo)
		liberr.ErrIsNil(ctx, err, "获取用户信息失败")
	})
	userTypes = strings.Split(userInfo.Type, ",")
	return
}

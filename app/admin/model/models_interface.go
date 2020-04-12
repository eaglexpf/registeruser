// model 包
package model

import (
	"context"
	"registeruser/app/admin/entity/dao"
)

// 后台用户接口
type AdminUserModel interface {
	FindUserByUUID(context.Context, string) (*dao.AdminUser, error)
	FindUserByUsername(context.Context, string) (*dao.AdminUser, error)
	FindUserByEmail(context.Context, string) (*dao.AdminUser, error)
	FindUserByMobile(context.Context, string) (*dao.AdminUser, error)
	InsertUser(context.Context, *dao.AdminUser) error
	UpdateUserInfoByUUID(context.Context, *dao.AdminUser) error
	UpdateUserPwdByUUID(context.Context, *dao.AdminUser) error
}

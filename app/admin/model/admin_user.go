package model

import (
	"context"
	"registeruser/app/admin/entity"
)

type AdminUserModel interface {
	FindUserByUUID(context.Context, string) (*entity.AdminUser, error)
	FindUserByUsername(context.Context, string) (*entity.AdminUser, error)
	FindUserByEmail(context.Context, string) (*entity.AdminUser, error)
	FindUserByMobile(context.Context, string) (*entity.AdminUser, error)
	InsertUser(context.Context, *entity.AdminUser) error
	UpdateUserInfoByUUID(context.Context, *entity.AdminUser) error
	UpdateUserPwdByUUID(context.Context, *entity.AdminUser) error
}

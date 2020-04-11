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
	InsertUser(ctx context.Context, user *entity.AdminUser) error
}

package model

import (
	"context"
	"registeruser/app/admin/entity"
)

type AdminUserModel interface {
	FindUserByUsername(ctx context.Context, username string) (*entity.AdminUser, error)
	InsertUser(ctx context.Context, user *entity.AdminUser) error
}

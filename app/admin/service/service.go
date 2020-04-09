package service

import (
	"context"
	"registeruser/app/admin/entity"
	"registeruser/app/admin/model"
	"registeruser/app/admin/model/mysql"
)

func NewService() *service {
	return &service{
		adminUserModel: mysql.NewAdminUserModel(),
	}
}

type service struct {
	adminUserModel model.AdminUserModel
}

func (this *service) FindUserByUsername(ctx context.Context, username string) (*entity.AdminUser, error) {
	return this.adminUserModel.FindUserByUsername(ctx, username)
}

func (this *service) RegisterUser(ctx context.Context, user *entity.AdminUser) error {
	return this.adminUserModel.InsertUser(ctx, user)
}

package service

import (
	"registeruser/app/admin/model"
	"registeruser/app/admin/model/mysql"
)

func NewService() *Service {
	return &Service{
		adminUserModel: mysql.NewAdminUserModel(),
	}
}

type Service struct {
	adminUserModel model.AdminUserModel
}

// service 服务包
package service

import (
	"registeruser/app/admin/model"
	"registeruser/app/admin/model/mysql"
)

// 返回一个新的service服务
func NewService() *Service {
	return &Service{
		adminUserModel: mysql.NewAdminUserModel(),
	}
}

// service服务类型
type Service struct {
	adminUserModel model.AdminUserModel
}

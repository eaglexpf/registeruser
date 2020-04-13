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
		adminRoleModel: mysql.NewAdminRoleModel(),
		adminApiModel:  mysql.NewAdminApiModel(),
	}
}

// service服务类型
type Service struct {
	adminUserModel model.AdminUserModel
	adminRoleModel model.AdminRoleModel
	adminApiModel  model.AdminApiModel
}

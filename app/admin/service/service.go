// service 服务包
package service

import (
	"registeruser/app/admin/entity/response"
	"registeruser/app/admin/model"
	"registeruser/app/admin/model/mysql"
)

// 返回一个新的service服务
func NewService() *Service {
	return &Service{
		adminUserModel:      mysql.NewAdminUserModel(),
		adminRoleModel:      mysql.NewAdminRoleModel(),
		adminApiModel:       mysql.NewAdminApiModel(),
		adminPermissionMode: mysql.NewAdminPermissionModel(),
	}
}

//func NewService() ServiceInter {
//	return &Service{
//		adminUserModel: mysql.NewAdminUserModel(),
//		adminRoleModel: mysql.NewAdminRoleModel(),
//		adminApiModel:  mysql.NewAdminApiModel(),
//	}
//}

type ServiceInter interface {
	Page(count, page, page_size, page_count int64, list interface{}) *response.ResponsePage
}

// service服务类型
type Service struct {
	adminUserModel      model.AdminUserModel
	adminRoleModel      model.AdminRoleModel
	adminApiModel       model.AdminApiModel
	adminPermissionMode model.AdminPermissionModel
}

func (s *Service) Page(count, page, page_size, page_count int64, list interface{}) *response.ResponsePage {
	responseData := &response.ResponsePage{
		Meta: &response.ResponsePageMeta{
			Count:     count,
			Page:      page,
			PageSize:  page_size,
			PageCount: page_count,
		},
		List: list,
	}
	return responseData
}

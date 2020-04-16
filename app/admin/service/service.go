// service 服务包
package service

import (
	"context"
	"registeruser/app/admin/entity/dao"
	"registeruser/app/admin/entity/response"
	"registeruser/app/admin/model"
	"registeruser/app/admin/model/mysql"
	"registeruser/app/admin/model/redis"
	"registeruser/util/common"
)

// 返回一个新的service服务
func NewService() *Service {
	return &Service{
		adminUserModel:       mysql.NewAdminUserModel(),
		adminUserCache:       redis.NewAdminUserCache(),
		adminRoleModel:       mysql.NewAdminRoleModel(),
		adminApiModel:        mysql.NewAdminApiModel(),
		adminServiceModel:    mysql.NewAdminServiceModel(),
		adminPermissionModel: mysql.NewAdminPermissionModel(),
	}
}

//func NewService() ServiceInter {
//	return &Service{
//		adminUserModel: mysql.NewAdminUserModel(),
//		adminRoleModel: mysql.NewAdminRoleModel(),
//		adminApiModel:  mysql.NewAdminApiModel(),
//	}
//}

// service服务类型
type Service struct {
	adminUserModel       model.AdminUserModel
	adminUserCache       model.AdminUserCache
	adminRoleModel       model.AdminRoleModel
	adminApiModel        model.AdminApiModel
	adminServiceModel    model.AdminServiceModel
	adminPermissionModel model.AdminPermissionModel
}

func (s *Service) AuthCheck(ctx context.Context, name string) (err error) {
	// 公共的不需要授权的服务
	service_no_auth := []string{"admin.user.login"}
	for _, v := range service_no_auth {
		if v == name {
			return
		}
	}
	user, ok := ctx.Value("adminUser").(*dao.AdminUser)
	if !ok {
		// jwt验证失败
		err = response.ERROR_AUTH
		return
	}
	// 超级管理员，不需要授权
	user_no_auth := []string{"admin"}
	for _, v := range user_no_auth {
		if v == user.UserName {
			return
		}
	}
	// 检查缓存是否存在
	if ok, err = s.adminUserCache.FindNameByUserUUID(ctx, user.UUID, name); err == nil {
		if !ok {
			// 缓存中没有查询到权限
			err = response.ERROR_AUTH
		}
		return
	}
	// 从数据库中获取
	var ids []int64
	ids = s.adminPermissionModel.FindPermissionListByUserID(ctx, user.ID, ids)
	if len(ids) == 0 {
		// 用户数据库中没有分配权限
		err = response.ERROR_AUTH
		return
	}
	ids = common.SliceDuplicateInt64(ids)
	serviceList, err := s.adminServiceModel.FindByIds(ctx, ids)
	if err != nil {
		// 根据权限id查询服务列表失败
		err = response.ERROR_AUTH
		return
	}
	if len(serviceList) == 0 {
		// 未查询到用户数据库中分配到的权限
		err = response.ERROR_AUTH
		return
	}
	result := make([]string, 0)
	for _, value := range serviceList {
		if value.Status == 1 && value.ExpireAt == 0 {
			result = append(result, value.Name)
		}
	}
	result = common.SliceDuplicateString(result)
	// 将结果放入缓存中
	go s.adminUserCache.SetUserServiceByUUID(ctx, user.UUID, result)
	for _, value := range result {
		if value == name {
			return
		}
	}
	// 用户数据库中没有查到该权限
	err = response.ERROR_AUTH
	return
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

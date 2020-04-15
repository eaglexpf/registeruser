package service

import (
	"context"
	"registeruser/app/admin/entity/dao"
	"registeruser/app/admin/model"
	"registeruser/app/admin/model/mysql"
)

// 权限验证

// 登录缓存用户对应的所有api权限
// 1.获取当前用户
// 2.获取当前api
// 3.查询api在权限表中所有的上级
// 4.判断上级user类型权限中是否有当前用户
// 5.
type ServiceAuth interface {
	FindUserByUUID(ctx context.Context, uuid string) (*dao.AdminUser, error)
	FindAllPermissionByUserID(ctx context.Context, user_id int64) []int64
}

func NewServiceAuth() ServiceAuth {
	return &serviceAuthStruct{
		adminUserModel:       mysql.NewAdminUserModel(),
		adminPermissionModel: mysql.NewAdminPermissionModel(),
	}
}

type serviceAuthStruct struct {
	adminUserModel       model.AdminUserModel
	adminPermissionModel model.AdminPermissionModel
}

func (s *serviceAuthStruct) FindUserByUUID(ctx context.Context, uuid string) (*dao.AdminUser, error) {
	return s.adminUserModel.FindUserByUUID(ctx, uuid)
}

func (s *serviceAuthStruct) FindAllPermissionByUserID(ctx context.Context, user_id int64) []int64 {
	var ids []int64
	ids = s.adminPermissionModel.FindPermissionListByUserID(ctx, user_id, ids)
	return ids
}

func role(id int64) {
	// 角色的递归查询

	// 从上到下
	//sql := `select * from admin_permission where parent_id=? and  parent_type="role" and children_type="role"`

}

// 将用户所有的api权限缓存
// 缓存api对应的用户
// 缓存角色对应的用户

// 删除用户时联动删除权限表中的对应的user权限					删除用户缓存
// 删除api时联动删除权限表中对应的api权限						删除对应用户中的缓存
// 删除角色时联动删除权限表中对应的role权限【父子结点都有】		删除对应用户中的缓存

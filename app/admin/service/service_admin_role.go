package service

import (
	"context"
	"registeruser/app/admin/entity/dao"
	"registeruser/app/admin/entity/request"
)

// 查询角色列表
func (s *Service) FindRoleList(ctx context.Context, page, page_size int64) (data []*dao.AdminRole, err error) {
	offset := (page - 1) * page_size
	data, err = s.adminRoleModel.FindRoleList(ctx, offset, page_size)
	return
}

// 根据角色id查询角色信息
func (s *Service) FindAdminRoleByID(ctx context.Context, id int64) (adminRole *dao.AdminRole, err error) {
	adminRole, err = s.adminRoleModel.FindRoleByID(ctx, id)
	return
}

// 添加一个角色
func (s *Service) RegisterAdminRole(ctx context.Context, request *request.RequestAdminRoleRegister) (adminRole *dao.AdminRole, err error) {
	adminRole = &dao.AdminRole{
		Name:        request.Name,
		Description: request.Description,
	}
	err = s.adminRoleModel.InsertRole(ctx, adminRole)
	return
}

// 修改角色名称描述等信息
func (s *Service) UpdateRoleByID(ctx context.Context, request *request.RequestAdminRoleUpdate) (admin_role *dao.AdminRole, err error) {
	admin_role, err = s.FindAdminRoleByID(ctx, request.ID)
	if err != nil {
		return
	}
	admin_role.Name = request.Name
	admin_role.Description = request.Description
	err = s.adminRoleModel.UpdateRoleByID(ctx, admin_role)
	return
}

// 删除一个角色
func (s *Service) DeleteAdminRoleByID(ctx context.Context, id int64) error {
	return s.adminRoleModel.DeleteRoleByID(ctx, id)
}

package service

import (
	"context"
	"registeruser/app/admin/entity/dao"
	"registeruser/app/admin/entity/request"
	"registeruser/app/admin/entity/response"
	"registeruser/conf/log"
)

// 查询角色列表
func (s *Service) FindRoleList(ctx context.Context, page, page_size int64) (data []dao.AdminRole, err error) {
	err = s.AuthCheck(ctx, "admin.role.list")
	if err != nil {
		return
	}
	if page < 1 {
		page = 1
	}
	if page_size < 1 {
		page_size = 20
	}
	offset := (page - 1) * page_size
	data, err = s.adminRoleModel.FindRoleList(ctx, offset, page_size)
	if err != nil {
		// 角色列表查询失败
		log.Errorf("角色列表查询失败，失败原因：%s", err.Error())
		err = response.ERROR_ROLE_LIST
	}
	return
}

// 根据角色id查询角色信息
func (s *Service) FindAdminRoleByID(ctx context.Context, id int64) (adminRole *dao.AdminRole, err error) {
	err = s.AuthCheck(ctx, "admin.role.find_by_id")
	if err != nil {
		return
	}
	adminRole, err = s.adminRoleModel.FindRoleByID(ctx, id)
	if err != nil {
		// 无效的角色
		err = response.ERROR_ROLE_FIND_BY_ID
	}
	return
}

// 添加一个角色
func (s *Service) RegisterAdminRole(ctx context.Context, request *request.RequestAdminRoleRegister) (adminRole *dao.AdminRole, err error) {
	err = s.AuthCheck(ctx, "admin.role.register")
	if err != nil {
		return
	}
	adminRole = &dao.AdminRole{
		Name:        request.Name,
		Description: request.Description,
	}
	err = s.adminRoleModel.InsertRole(ctx, adminRole)
	if err != nil {
		// 角色添加失败
		log.Errorf("角色添加失败，失败原因：%s", err.Error())
		err = response.ERROR_ROLE_REGISTER
	}
	return
}

// 修改角色名称描述等信息
func (s *Service) UpdateRoleByID(ctx context.Context, request *request.RequestAdminRoleUpdate) (admin_role *dao.AdminRole, err error) {
	err = s.AuthCheck(ctx, "admin.role.update")
	if err != nil {
		return
	}
	admin_role, err = s.FindAdminRoleByID(ctx, request.ID)
	if err != nil {
		return
	}
	admin_role.Name = request.Name
	admin_role.Description = request.Description
	err = s.adminRoleModel.UpdateRoleByID(ctx, admin_role)
	if err != nil {
		// 角色修改失败
		log.Errorf("角色修改失败，失败原因：%s", err.Error())
		err = response.ERROR_ROLE_UPDATE
	}
	return
}

// 删除一个角色
func (s *Service) DeleteAdminRoleByID(ctx context.Context, id int64) (err error) {
	err = s.AuthCheck(ctx, "admin.role.delete")
	if err != nil {
		return
	}
	_, err = s.FindAdminRoleByID(ctx, id)
	if err != nil {
		return
	}
	err = s.adminRoleModel.DeleteRoleByID(ctx, id)
	if err != nil {
		// 角色删除失败
		log.Errorf("角色删除失败，失败原因：%s", err.Error())
		err = response.ERROR_ROLE_DELETE
	}
	return
}

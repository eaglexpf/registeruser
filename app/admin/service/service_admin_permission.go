package service

import (
	"context"
	"registeruser/app/admin/entity/dao"
	"registeruser/app/admin/entity/request"
	"registeruser/app/admin/entity/response"
	"registeruser/conf/log"
)

// 检查parent_type和children_type是否正确
func (s *Service) checkDaoAdminPermission(data *dao.AdminPermission) error {
	if data.ParentType != "" {
		if data.ParentType != "user" && data.ParentType != "role" {
			return response.ERROR_PERMISSION_DAO_PARENT_TYPE
		}
	}
	if data.ChildrenType != "" {
		if data.ChildrenType != "role" && data.ChildrenType != "api" && data.ChildrenType != "service" {
			return response.ERROR_PERMISSION_DAO_CHILDREN_TYPE
		}
	}
	return nil
}

// 查询权限
func (s *Service) PermissionSearch(ctx context.Context, searchRequest *request.RequestPermissionSearch) (responseData *response.ResponsePage, err error) {
	err = s.AuthCheck(ctx, "admin.permission.search")
	if err != nil {
		return
	}
	if searchRequest.Page == 0 {
		searchRequest.Page = 1
	}
	if searchRequest.PageSize == 0 {
		searchRequest.PageSize = 20
	}
	offset := (searchRequest.Page - 1) * searchRequest.PageSize
	search := &dao.AdminPermission{
		ParentID:     searchRequest.ParentID,
		ParentType:   searchRequest.ParentType,
		ChildrenID:   searchRequest.ChildrenID,
		ChildrenType: searchRequest.ChildrenType,
	}
	// 检查数据是否正确
	err = s.checkDaoAdminPermission(search)
	if err != nil {
		return
	}

	list, err := s.adminPermissionModel.Search(ctx, search, offset, searchRequest.PageSize)
	if err != nil {
		// 权限分配列表查询失败
		log.Errorf("权限分配列表查询失败，失败原因：%s", err.Error())
		err = response.ERROR_PERMISSION_LIST
	}
	count := s.adminPermissionModel.SearchCount(ctx, search)
	responseData = s.Page(count, searchRequest.Page, searchRequest.PageSize, int64(len(list)), list)

	return
}

// 注册权限
func (s *Service) PermissionRegister(ctx context.Context, insertRequest *request.RequestPermissionRegister) (err error) {
	err = s.AuthCheck(ctx, "admin.permission.register")
	if err != nil {
		return
	}
	insert_data := &dao.AdminPermission{
		ParentID:     insertRequest.ParentID,
		ParentType:   insertRequest.ParentType,
		ChildrenID:   insertRequest.ChildrenID,
		ChildrenType: insertRequest.ChildrenType,
	}
	// 检查数据是否正确
	err = s.checkDaoAdminPermission(insert_data)
	if err != nil {
		return
	}
	err = s.adminPermissionModel.Register(ctx, insert_data)
	if err != nil {
		// 权限分配添加失败
		log.Errorf("权限分配添加失败，失败原因：%s", err.Error())
		err = response.ERROR_PERMISSION_REGISTER
	}
	return err
}

// 删除权限
func (s *Service) PermissionDelete(ctx context.Context, deleteRequest *request.RequestPermissionDelete) (err error) {
	err = s.AuthCheck(ctx, "admin.permission.delete")
	if err != nil {
		return
	}
	if deleteRequest.ParentID > 0 {
		if deleteRequest.ParentType == "" {
			// 父权限类型不能为空
			err = response.ERROR_PERMISSION_NONE_PARENT_TYPE
		}
	}
	if deleteRequest.ParentType != "" {
		if deleteRequest.ParentID == 0 {
			// 父权限id不能为空
			err = response.ERROR_PERMISSION_NONE_PARENT_ID
		}
	}
	if deleteRequest.ChildrenID > 0 {
		if deleteRequest.ChildrenType == "" {
			// 子权限类型不能为空
			err = response.ERROR_PERMISSION_NONE_CHILDREN_TYPE
		}
	}
	if deleteRequest.ChildrenType != "" {
		if deleteRequest.ChildrenID == 0 {
			// 子权限id不能为空
			err = response.ERROR_PERMISSION_NONE_CHILDREN_ID
		}
	}
	if deleteRequest.ParentType == "" && deleteRequest.ParentID == 0 && deleteRequest.ChildrenType == "" && deleteRequest.ChildrenID == 0 {
		// 不能删除空节点
		err = response.ERROR_PERMISSION_NONE
	}
	if err != nil {
		return
	}
	data := &dao.AdminPermission{
		ParentID:     deleteRequest.ParentID,
		ParentType:   deleteRequest.ParentType,
		ChildrenID:   deleteRequest.ChildrenID,
		ChildrenType: deleteRequest.ChildrenType,
	}
	// 检查数据是否正确
	err = s.checkDaoAdminPermission(data)
	if err != nil {
		return
	}
	err = s.adminPermissionModel.DeleteInfo(ctx, data)
	if err != nil {
		// 权限分配删除失败
		log.Errorf("权限分配删除失败，失败原因：%s", err.Error())
		err = response.ERROR_PERMISSION_DELETE
	}
	return
}

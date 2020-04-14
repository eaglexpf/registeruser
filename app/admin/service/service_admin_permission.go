package service

import (
	"context"
	"errors"
	"registeruser/app/admin/entity/dao"
	"registeruser/app/admin/entity/request"
	"registeruser/app/admin/entity/response"
)

// 查询权限
func (s *Service) PermissionSearch(ctx context.Context, searchRequest *request.RequestPermissionSearch) (responseData *response.ResponsePage, err error) {
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
	list, err := s.adminPermissionMode.Search(ctx, search, offset, searchRequest.PageSize)
	if err != nil {
		return
	}
	count := s.adminPermissionMode.SearchCount(ctx, search)
	responseData = s.Page(count, searchRequest.Page, searchRequest.PageSize, int64(len(list)), list)
	return
}

// 注册权限
func (s *Service) PermissionRegister(ctx context.Context, insertRequest *request.RequestPermissionRegister) error {
	insert_data := &dao.AdminPermission{
		ParentID:     insertRequest.ParentID,
		ParentType:   insertRequest.ParentType,
		ChildrenID:   insertRequest.ChildrenID,
		ChildrenType: insertRequest.ChildrenType,
	}
	err := s.adminPermissionMode.Register(ctx, insert_data)
	return err
}

// 删除权限
func (s *Service) PermissionDelete(ctx context.Context, deleteRequest *request.RequestPermissionDelete) (err error) {
	switch true {
	case deleteRequest.ParentID > 0:
		if deleteRequest.ParentType == "" {
			err = errors.New("父权限类型不能为空")
		}
	case deleteRequest.ParentType != "":
		if deleteRequest.ParentID == 0 {
			err = errors.New("父权限id不能为空")
		}
	case deleteRequest.ChildrenID > 0:
		if deleteRequest.ChildrenType != "" {
			err = errors.New("子权限类型不能为空")
		}
	case deleteRequest.ChildrenType != "":
		if deleteRequest.ChildrenID == 0 {
			err = errors.New("子权限id不能为空")
		}
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
	err = s.adminPermissionMode.DeleteInfo(ctx, data)
	return
}

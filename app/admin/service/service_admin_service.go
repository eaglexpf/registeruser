package service

import (
	"context"
	"registeruser/app/admin/entity/dao"
	"registeruser/app/admin/entity/request"
	"registeruser/app/admin/entity/response"
	"registeruser/conf/log"
)

func (s *Service) ServiceFindByID(ctx context.Context, id int64) (data *dao.AdminService, err error) {
	err = s.AuthCheck(ctx, "admin.service.find_by_id")
	if err != nil {
		return
	}
	data, err = s.adminServiceModel.FindByID(ctx, id)
	if err != nil {
		// 无效的服务
		err = response.ERROR_SERVICE_FIND_BY_ID
	}
	return
}

func (s *Service) ServiceSearch(ctx context.Context, data *request.RequestServiceSearch) (responseData *response.ResponsePage, err error) {
	err = s.AuthCheck(ctx, "admin.service.search")
	if err != nil {
		return
	}
	if data.Page == 0 {
		data.Page = 1
	}
	if data.PageSize == 0 {
		data.PageSize = 20
	}
	offset := (data.Page - 1) * data.PageSize
	search := &dao.AdminServiceSearch{
		Name:   data.Name,
		Alias:  data.Alias,
		Status: data.Status,
	}
	list, err := s.adminServiceModel.Search(ctx, search, offset, data.PageSize)
	if err != nil {
		// 服务列表查询失败
		log.Errorf("服务列表查询失败，失败原因：%s", err.Error())
		err = response.ERROR_SERVICE_LIST
		return
	}
	count := s.adminServiceModel.SearchCount(ctx, search)
	responseData = s.Page(count, data.Page, data.PageSize, int64(len(list)), &list)
	return
}

func (s *Service) ServiceRegister(ctx context.Context, data *request.RequestServiceRegister) (err error) {
	err = s.AuthCheck(ctx, "admin.service.register")
	if err != nil {
		return
	}
	_, err = s.adminServiceModel.FindByName(ctx, data.Name)
	if err == nil {
		// 服务名称已存在
		err = response.ERROR_SERVICE_UNIQUE_NAME
		return
	}
	daoData := &dao.AdminService{
		Name:        data.Name,
		Alias:       data.Alias,
		Description: data.Description,
		Status:      data.Status,
		ExpireAt:    data.ExpireAt,
	}
	err = s.adminServiceModel.Register(ctx, daoData)
	if err != nil {
		// 服务添加失败
		log.Errorf("服务添加失败，失败原因：%s", err.Error())
		err = response.ERROR_SERVICE_REGISTER
	}
	return err
}

func (s *Service) ServiceUpdateByID(ctx context.Context, data *request.RequestServiceUpdate, id int64) (err error) {
	err = s.AuthCheck(ctx, "admin.service.update")
	if err != nil {
		return
	}
	dataData, err := s.adminServiceModel.FindByID(ctx, id)
	if err != nil {
		// 无效的服务
		err = response.ERROR_SERVICE_FIND_BY_ID
		return
	}
	dataData.Name = data.Name
	dataData.Alias = data.Alias
	dataData.Description = data.Description
	dataData.Status = data.Status
	dataData.ExpireAt = data.ExpireAt
	err = s.adminServiceModel.UpdateByID(ctx, dataData)
	if err != nil {
		// 服务修改失败
		log.Errorf("服务修改失败，失败原因：%s", err.Error())
		err = response.ERROR_SERVICE_UPDATE
	}
	return
}

func (s *Service) ServiceDeleteByID(ctx context.Context, id int64) (err error) {
	err = s.AuthCheck(ctx, "admin.service.delete_by_id")
	if err != nil {
		return
	}
	_, err = s.adminServiceModel.FindByID(ctx, id)
	if err != nil {
		// 无效的服务
		err = response.ERROR_SERVICE_FIND_BY_ID
		return
	}
	err = s.adminServiceModel.DeleteByID(ctx, id)
	if err != nil {
		// 服务删除失败
		log.Errorf("服务根据服务id删除失败，失败原因：%s", err.Error())
		err = response.ERROR_SERVICE_DELETE_BY_ID
	}
	return
}

func (s *Service) ServiceDeleteByName(ctx context.Context, name string) (err error) {
	err = s.AuthCheck(ctx, "admin.service.delete_by_name")
	if err != nil {
		return
	}
	_, err = s.adminServiceModel.FindByName(ctx, name)
	if err != nil {
		// 无效的服务
		err = response.ERROR_SERVICE_FIND_BY_ID
		return
	}
	err = s.adminServiceModel.DeleteByName(ctx, name)
	if err != nil {
		// 服务删除失败
		log.Errorf("服务根据服务名删除失败，失败原因：%s", err.Error())
		err = response.ERROR_SERVICE_DELETE_BY_NAME
	}
	return
}

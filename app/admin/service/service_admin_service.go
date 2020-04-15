package service

import (
	"context"
	"errors"
	"registeruser/app/admin/entity/dao"
	"registeruser/app/admin/entity/request"
	"registeruser/app/admin/entity/response"
)

func (s *Service) ServiceFindByID(ctx context.Context, id int64) (data *dao.AdminService, err error) {
	data, err = s.adminServiceModel.FindByID(ctx, id)
	return
}

func (s *Service) ServiceSearch(ctx context.Context, data *request.RequestServiceSearch) (responseData *response.ResponsePage, err error) {
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
		return
	}
	count := s.adminServiceModel.SearchCount(ctx, search)
	responseData = s.Page(count, data.Page, data.PageSize, int64(len(list)), &list)
	return
}

func (s *Service) ServiceRegister(ctx context.Context, data *request.RequestServiceRegister) error {
	_, err := s.adminServiceModel.FindByName(ctx, data.Name)
	if err == nil {
		return errors.New("服务名称已存在")
	}
	daoData := &dao.AdminService{
		Name:        data.Name,
		Alias:       data.Alias,
		Description: data.Description,
		Status:      data.Status,
		ExpireAt:    data.ExpireAt,
	}
	err = s.adminServiceModel.Register(ctx, daoData)
	return err
}

func (s *Service) ServiceUpdateByID(ctx context.Context, data *request.RequestServiceUpdate, id int64) error {
	dataData, err := s.adminServiceModel.FindByID(ctx, id)
	if err != nil {
		return err
	}
	dataData.Name = data.Name
	dataData.Alias = data.Alias
	dataData.Description = data.Description
	dataData.Status = data.Status
	dataData.ExpireAt = data.ExpireAt
	err = s.adminServiceModel.UpdateByID(ctx, dataData)
	return err
}

func (s *Service) ServiceDeleteByID(ctx context.Context, id int64) error {
	_, err := s.adminServiceModel.FindByID(ctx, id)
	if err != nil {
		return err
	}
	err = s.adminServiceModel.DeleteByID(ctx, id)
	return err
}

func (s *Service) ServiceDeleteByName(ctx context.Context, name string) error {
	_, err := s.adminServiceModel.FindByName(ctx, name)
	if err != nil {
		return err
	}
	err = s.adminServiceModel.DeleteByName(ctx, name)
	return err
}

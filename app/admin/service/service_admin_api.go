package service

import (
	"context"
	"registeruser/app/admin/entity/dao"
	"registeruser/app/admin/entity/request"
	"registeruser/app/admin/entity/response"
)

func (s *Service) ApiFindAll(ctx context.Context, page, page_size int64) (responseData *response.ResponsePage, err error) {
	offset := (page - 1) * page_size
	list, count, err := s.adminApiModel.FindAll(ctx, offset, page_size)
	if err != nil {
		return
	}
	responseData = s.Page(count, page, page_size, int64(len(list)), &list)
	return
}

func (s *Service) ApiFindByID(ctx context.Context, id int64) (data *dao.AdminApi, err error) {
	data, err = s.adminApiModel.FindByID(ctx, id)
	return
}

func (s *Service) ApiSearch(ctx context.Context, method, path string, page, page_size int64) (responseData *response.ResponsePage, err error) {
	offset := (page - 1) * page_size
	list, count, err := s.adminApiModel.Search(ctx, method, path, offset, page_size)
	if err != nil {
		return
	}
	responseData = s.Page(count, page, page_size, int64(len(list)), &list)
	return
}

func (s *Service) ApiRegister(ctx context.Context, request *request.RequestApiRegister) error {
	data := &dao.AdminApi{
		Method:      request.Method,
		Path:        request.Path,
		Description: request.Description,
	}
	err := s.adminApiModel.Register(ctx, data)
	return err
}

func (s *Service) ApiUpdateByID(ctx context.Context, request *request.RequestApiUpdate, id int64) error {
	data := &dao.AdminApi{
		ID:          id,
		Method:      request.Method,
		Path:        request.Path,
		Description: request.Description,
	}
	err := s.adminApiModel.UpdateByID(ctx, data)
	return err
}

func (s *Service) ApiDeleteByID(ctx context.Context, id int64) error {
	err := s.adminApiModel.DeleteByID(ctx, id)
	return err
}

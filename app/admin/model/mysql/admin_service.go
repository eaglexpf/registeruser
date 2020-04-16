package mysql

import (
	"context"
	"errors"
	"fmt"
	"registeruser/app/admin/entity/dao"
	"registeruser/app/admin/model"
	"registeruser/conf/global"
	"registeruser/util/sql_util"
	"strings"
	"time"
)

const (
	QUERY_SERVICE_FIND_BY_ID   = `select * from admin_service where id=?`
	QUERY_SERVICE_FIND_BY_IDS  = `select * from admin_service where id in (%s)`
	QUERY_SERVICE_FIND_BY_NAME = `select * from admin_service where name=?`

	QUERY_SERVICE_SEARCH_ALL                  = `select * from admin_service limit ?,?`
	QUERY_SERVICE_SEARCH_BY_NAME              = `select * from admin_service where name like "%?%" limit ?,?`
	QUERY_SERVICE_SEARCH_BY_NAME_ALIAS        = `select * from admin_service where name like "%?%" and alias like "%?%" limit ?,?`
	QUERY_SERVICE_SEARCH_BY_NAME_STATUS       = `select * from admin_service where name like "%?%" and status=? limit ?,?`
	QUERY_SERVICE_SEARCH_BY_NAME_ALIAS_STATUS = `select * from admin_service where name like "%?%" and alias like "%?%" and status=? limit ?,?`
	QUERY_SERVICE_SEARCH_BY_ALIAS             = `select * from admin_service where alias like "%?%" limit ?,?`
	QUERY_SERVICE_SEARCH_BY_ALIAS_STATUS      = `select * from admin_service where alias like "%?%" and status=? limit ?,?`
	QUERY_SERVICE_SEARCH_BY_STATUS            = `select * from admin_service where status=? limit ?,?`

	QUERY_SERVICE_SEARCH_ALL_COUNT                   = `select count(*) as count from admin_service`
	QUERY_SERVICE_SEARCH_BY_NAME_COUNT               = `select count(*) as count from admin_service where name like "%?%"`
	QUERY_SERVICE_SEARCH_BY_NAME_ALIAS_COUNT         = `select count(*) as count from admin_service where name like "%?%" and alias like "%?%"`
	QUERY_SERVICE_SEARCH_BY_NAME_STATUS_COUNT        = `select count(*) as count from admin_service where name like "%?%" and status=?`
	QUERY_SERVICE_SEARCH_BY_NAME__ALIAS_STATUS_COUNT = `select count(*) as count from admin_service where name like "%?%" and alias like "%?%" and status=?`
	QUERY_SERVICE_SEARCH_BY_ALIAS_COUNT              = `select count(*) as count from admin_service where alias like "%?%"`
	QUERY_SERVICE_SEARCH_BY_ALIAS_STATUS_COUNT       = `select count(*) as count from admin_service where alias like "%?%" and status=?`
	QUERY_SERVICE_SEARCH_BY_STATUS_COUNT             = `select count(*) as count from admin_service where status=?`

	QUERY_SERVICE_REGISTER = `insert into admin_service ( name,alias,description,status,expire_at,create_at ) values ( ?,?,?,?,?,? )`

	QUERY_SERVICE_UPDATE_BY_ID = `update admin_service set name=?,alias=?,description=?,status=?,expire_at=?,update_at=? where id=?`

	QUERY_SERVICE_DELETE_BY_ID   = `delete from admin_service where id=?`
	QUERY_SERVICE_DELETE_BY_NAME = `delete from admin_service where name=?`
)

func NewAdminServiceModel() model.AdminServiceModel {
	return &serviceModel{
		sql_util.NewSqlUtil(global.DB),
	}
}

type serviceModel struct {
	*sql_util.SqlUtil
}

func (s *serviceModel) checkSearch(data *dao.AdminServiceSearch) error {
	switch data.Status {
	case 0, 1, 2:
	default:
		return errors.New("服务的状态只能为0，1，2")
	}
	return nil
}
func (s *serviceModel) checkRegister(data *dao.AdminService) error {
	switch data.Status {
	case 0, 1, 2:
	default:
		return errors.New("服务的状态只能为0，1，2")
	}
	return nil
}

func (s *serviceModel) Search(ctx context.Context, search *dao.AdminServiceSearch, offset, limit int64) (result []dao.AdminService, err error) {
	if err = s.checkSearch(search); err != nil {
		return
	}
	switch true {
	case search.Name != "" && search.Alias != "" && search.Status > 0:
		err = s.Fetch(ctx, QUERY_SERVICE_SEARCH_BY_NAME_ALIAS_STATUS, &result, search.Name, search.Alias, search.Status, offset, limit)
	case search.Name != "" && search.Alias != "":
		err = s.Fetch(ctx, QUERY_SERVICE_SEARCH_BY_NAME_ALIAS, &result, search.Name, search.Alias, offset, limit)
	case search.Name != "" && search.Status > 0:
		err = s.Fetch(ctx, QUERY_SERVICE_SEARCH_BY_NAME_STATUS, &result, search.Name, search.Status, offset, limit)
	case search.Alias != "" && search.Status > 0:
		err = s.Fetch(ctx, QUERY_SERVICE_SEARCH_BY_ALIAS_STATUS, &result, search.Alias, search.Status, offset, limit)
	case search.Name != "":
		err = s.Fetch(ctx, QUERY_SERVICE_SEARCH_BY_NAME, &result, search.Name, offset, limit)
	case search.Alias != "":
		err = s.Fetch(ctx, QUERY_SERVICE_SEARCH_BY_ALIAS, &result, search.Alias, offset, limit)
	case search.Status > 0:
		err = s.Fetch(ctx, QUERY_SERVICE_SEARCH_BY_STATUS, &result, search.Status, offset, limit)
	default:
		err = s.Fetch(ctx, QUERY_SERVICE_SEARCH_ALL, &result, offset, limit)
	}
	fmt.Println(result)
	return
}

func (s *serviceModel) SearchCount(ctx context.Context, search *dao.AdminServiceSearch) int64 {
	var result map[string]interface{}
	var err error
	if err = s.checkSearch(search); err != nil {
		return 0
	}
	switch true {
	case search.Name != "" && search.Alias != "" && search.Status > 0:
		result, err = s.FetchMapRow(ctx, QUERY_SERVICE_SEARCH_BY_NAME__ALIAS_STATUS_COUNT, search.Name, search.Alias, search.Status)
	case search.Name != "" && search.Alias != "":
		result, err = s.FetchMapRow(ctx, QUERY_SERVICE_SEARCH_BY_NAME_ALIAS_COUNT, search.Name, search.Alias)
	case search.Name != "" && search.Status > 0:
		result, err = s.FetchMapRow(ctx, QUERY_SERVICE_SEARCH_BY_NAME_STATUS_COUNT, search.Name, search.Status)
	case search.Alias != "" && search.Status > 0:
		result, err = s.FetchMapRow(ctx, QUERY_SERVICE_SEARCH_BY_ALIAS_STATUS_COUNT, search.Alias, search.Status)
	case search.Name != "":
		result, err = s.FetchMapRow(ctx, QUERY_SERVICE_SEARCH_BY_NAME_COUNT, search.Name)
	case search.Alias != "":
		result, err = s.FetchMapRow(ctx, QUERY_SERVICE_SEARCH_BY_ALIAS_COUNT, search.Alias)
	case search.Status > 0:
		result, err = s.FetchMapRow(ctx, QUERY_SERVICE_SEARCH_BY_STATUS_COUNT, search.Status)
	default:
		result, err = s.FetchMapRow(ctx, QUERY_SERVICE_SEARCH_ALL_COUNT)
	}
	if err != nil {
		return 0
	}
	if _, ok := result["count"]; !ok {
		return 0
	}
	return result["count"].(int64)
}

func (s *serviceModel) FindByID(ctx context.Context, id int64) (*dao.AdminService, error) {
	var request dao.AdminService
	err := s.FetchRow(ctx, QUERY_SERVICE_FIND_BY_ID, &request, id)
	return &request, err
}

func (s *serviceModel) FindByName(ctx context.Context, name string) (*dao.AdminService, error) {
	var request dao.AdminService
	err := s.FetchRow(ctx, QUERY_SERVICE_FIND_BY_NAME, &request, name)
	return &request, err
}

func (s *serviceModel) FindByIds(ctx context.Context, ids []int64) (result []dao.AdminService, err error) {
	var temp_ids []interface{}
	var temp_where []string
	for _, v := range ids {
		temp_where = append(temp_where, "?")
		temp_ids = append(temp_ids, v)
	}
	temp_query := fmt.Sprintf(QUERY_SERVICE_FIND_BY_IDS, strings.Join(temp_where, ","))
	err = s.Fetch(ctx, temp_query, &result, temp_ids...)
	return
}

func (s *serviceModel) Register(ctx context.Context, data *dao.AdminService) error {
	if err := s.checkRegister(data); err != nil {
		return err
	}
	_, err := s.Insert(ctx, QUERY_SERVICE_REGISTER, data.Name, data.Alias, data.Description, data.Status, data.ExpireAt, time.Now().Unix())
	return err
}

func (s *serviceModel) UpdateByID(ctx context.Context, data *dao.AdminService) error {
	switch data.Status {
	case 1, 2:
	default:
		return errors.New("服务的状态只能为1，2")
	}
	affect, err := s.Update(ctx, QUERY_SERVICE_UPDATE_BY_ID, data.Name, data.Alias, data.Description, data.Status, data.ExpireAt, time.Now().Unix(), data.ID)
	if err != nil {
		return err
	}
	if affect != 1 {
		err = fmt.Errorf("修改了%d条数据", affect)
		return err
	}
	return nil
}

func (s *serviceModel) DeleteByID(ctx context.Context, id int64) error {
	affect, err := s.Delete(ctx, QUERY_SERVICE_DELETE_BY_ID, id)
	if err != nil {
		return err
	}
	if affect != 1 {
		err = fmt.Errorf("修改了%d条数据", affect)
		return err
	}
	return nil
}

func (s *serviceModel) DeleteByName(ctx context.Context, name string) error {
	affect, err := s.Delete(ctx, QUERY_SERVICE_DELETE_BY_NAME, name)
	if err != nil {
		return err
	}
	if affect != 1 {
		err = fmt.Errorf("修改了%d条数据", affect)
		return err
	}
	return nil
}

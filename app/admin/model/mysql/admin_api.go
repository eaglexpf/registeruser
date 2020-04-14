package mysql

import (
	"context"
	"fmt"
	"registeruser/app/admin/entity/dao"
	"registeruser/app/admin/model"
	"registeruser/conf/global"
	"registeruser/util/sql_util"
	"time"
)

const (
	QUERY_API_FIND_ALL                     = `select * from admin_api limit ?,?`
	QUERY_API_FIND_ALL_COUNT               = `select count(*) as count from admin_api`
	QUERY_API_FIND_BY_ID                   = `select * from admin_api where id=?`
	QUERY_API_INSERT                       = `insert into admin_api ( method,path,description,create_at ) values ( ?, ?, ?, ? )`
	QUERY_API_UPDATE_BY_ID                 = `update admin_api set method=?,path=?,description=?,update_at=? where id=?`
	QUERY_API_DELETE_BY_ID                 = `delete from admin_api where id=?`
	QUERY_API_SEARCH_METHOD                = `select * from admin_api where method=? limit ?,?`
	QUERY_API_SEARCH_PATH                  = `select * from admin_api where path like "%?%" limit ?,?`
	QUERY_API_SEARCH_METHOD_AND_PATH       = `select * from admin_api where method=? and path like "%?%" limit ?,?`
	QUERY_API_SEARCH_METHOD_COUNT          = `select count(*) as count from admin_api where method=? limit ?,?`
	QUERY_API_SEARCH_PATH_COUNT            = `select count(*) as count from admin_api where path like "%?%" limit ?,?`
	QUERY_API_SEARCH_METHOD_AND_PATH_COUNT = `select count(*) as count from admin_api where method=? and path like "%?%" limit ?,?`
)

// 初始化admin_api表的model服务
func NewAdminApiModel() model.AdminApiModel {
	return &api{
		sql_util.NewSqlUtil(global.DB),
	}
}

type api struct {
	*sql_util.SqlUtil
}

func (a *api) count(ctx context.Context, query string, args ...interface{}) int64 {
	result, err := a.FetchMapRow(ctx, query, args...)
	if err != nil {
		return 0
	}
	count, ok := result["count"]
	if !ok {
		return 0
	}
	count_int, ok := count.(int64)
	if !ok {
		return 0
	}
	return count_int
}

// 查询全部数据
func (a *api) FindAll(ctx context.Context, offset, limit int64) (list []dao.AdminApi, count int64, err error) {
	err = a.Fetch(ctx, QUERY_API_FIND_ALL, &list, offset, limit)
	count = a.count(ctx, QUERY_API_FIND_ALL_COUNT)
	return
}

// 根据id查询
func (a *api) FindByID(ctx context.Context, id int64) (*dao.AdminApi, error) {
	var request dao.AdminApi
	//request := new(dao.AdminApi)
	err := a.FetchRow(ctx, QUERY_API_FIND_BY_ID, &request, id)
	return &request, err
}

// search搜索
func (a *api) Search(ctx context.Context, method, path string, offset, limit int64) (list []dao.AdminApi, count int64, err error) {
	switch true {
	case method == "" && path != "":
		err = a.Fetch(ctx, QUERY_API_SEARCH_PATH, &list, path, offset, limit)
		count = a.count(ctx, QUERY_API_SEARCH_PATH_COUNT, path)
	case method != "" && path == "":
		err = a.Fetch(ctx, QUERY_API_SEARCH_METHOD, &list, method, offset, limit)
		count = a.count(ctx, QUERY_API_SEARCH_METHOD_COUNT, method)
	case method != "" && path != "":
		err = a.Fetch(ctx, QUERY_API_SEARCH_METHOD_AND_PATH, &list, method, path, offset, limit)
		count = a.count(ctx, QUERY_API_SEARCH_METHOD_AND_PATH_COUNT, method, path)
	default:
		err = a.Fetch(ctx, QUERY_API_FIND_ALL, &list, offset, limit)
		count = a.count(ctx, QUERY_API_FIND_ALL_COUNT)
	}
	return
}

// 注册一条新数据
func (a *api) Register(ctx context.Context, data *dao.AdminApi) error {
	last_id, err := a.Insert(ctx, QUERY_API_INSERT, data.Method, data.Path, data.Description, time.Now().Unix())
	if err != nil {
		return err
	}
	if last_id <= 0 {
		err = fmt.Errorf("插入id为：%d", last_id)
		return err
	}
	return nil
}

// 根据id修改
func (a *api) UpdateByID(ctx context.Context, data *dao.AdminApi) error {
	affect, err := a.Update(ctx, QUERY_API_UPDATE_BY_ID, data.Method, data.Path, data.Description, time.Now().Unix(), data.ID)
	if err != nil {
		return err
	}
	if affect != 1 {
		err = fmt.Errorf("修改了%d条数据", affect)
		return err
	}
	return nil
}

// 根据id删除
func (a *api) DeleteByID(ctx context.Context, id int64) error {
	affect, err := a.Delete(ctx, QUERY_API_DELETE_BY_ID, id)
	if err != nil {
		return err
	}
	if affect != 1 {
		err = fmt.Errorf("修改了%d条数据", affect)
		return err
	}
	return nil
}

package mysql

import (
	"context"
	"errors"
	"fmt"
	"registeruser/app/admin/entity/dao"
	"registeruser/app/admin/model"
	"registeruser/conf/global"
	"registeruser/conf/log"
	"registeruser/util/sql_util"
	"time"
)

const (
	QUERY_FIND_ROLE_LIST    = `select id,name,description,create_at,update_at from admin_role limit ?,? `
	QUERY_FIND_ROLE_BY_ID   = `select id,name,description,create_at,update_at from admin_role where id=?`
	QUERY_INSERT_ROLE       = `insert into admin_role ( name, description, create_at ) values( ?, ?, ? )`
	QUERY_UPDATE_ROLE_BY_ID = `update admin_role set name=?, description=?, update_at=? where id=?`
	QUERY_DELETE_ROLE_BY_ID = `delete from admin_role where id=?`
)

func NewAdminRoleModel() model.AdminRoleModel {
	return &role{}
}

type role struct {
}

func (r *role) fetch(ctx context.Context, query string, args ...interface{}) ([]*dao.AdminRole, error) {
	rows, err := global.DB.QueryContext(ctx, query, args...)
	if err != nil {
		log.Error("admin_role fetch query error: ", err)
		return nil, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Error("admin_role fetch close error: ", err)
		}
	}()

	result := make([]*dao.AdminRole, 0)
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		row := new(dao.AdminRole)

		addrs, err := sql_util.AddrsEncode(row, columns)
		if err != nil {
			return nil, err
		}

		err = rows.Scan(addrs...)
		if err != nil {
			return nil, err
		}

		result = append(result, row)
	}
	return result, nil
}

func (r *role) fetchRow(ctx context.Context, query string, args ...interface{}) (*dao.AdminRole, error) {
	list, err := r.fetch(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	if len(list) <= 0 {
		return nil, errors.New("未查询到数据")
	}
	return list[0], nil
}

func (r *role) update(ctx context.Context, query string, args ...interface{}) error {
	fmt.Println(query, args)
	stmt, err := global.DB.PrepareContext(ctx, query)
	if err != nil {
		log.Error("admin_role update error:", err)
		return err
	}

	res, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		return err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affect != 1 {
		return fmt.Errorf("修改了%d条数据", affect)
	}
	return nil
}

func (r *role) delete(ctx context.Context, query string, args ...interface{}) error {
	stmt, err := global.DB.PrepareContext(ctx, query)
	if err != nil {
		log.Error("admin_role delete error:", err)
		return err
	}

	res, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		return err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affect != 1 {
		return fmt.Errorf("删除了%d条数据", affect)
	}
	return nil
}

func (r *role) FindRoleList(ctx context.Context, offset, limit int64) ([]*dao.AdminRole, error) {
	return r.fetch(ctx, QUERY_FIND_ROLE_LIST, offset, limit)
}

func (r *role) FindRoleByID(ctx context.Context, id int64) (adminRole *dao.AdminRole, err error) {
	adminRole, err = r.fetchRow(ctx, QUERY_FIND_ROLE_BY_ID, id)
	return
}

func (r *role) InsertRole(ctx context.Context, data *dao.AdminRole) error {
	stmt, err := global.DB.PrepareContext(ctx, QUERY_INSERT_ROLE)
	if err != nil {
		log.Error("admin_role insert prepare err: ", err)
		return err
	}

	res, err := stmt.ExecContext(ctx, data.Name, data.Description, time.Now().Unix())
	if err != nil {
		log.Error("admin_role insert error:", err)
		return err
	}

	data.ID, err = res.LastInsertId()
	if err != nil {
		log.Error("admin_role get insert id error:", err)
		return err
	}
	return nil
}

func (r *role) UpdateRoleByID(ctx context.Context, data *dao.AdminRole) error {
	return r.update(ctx, QUERY_UPDATE_ROLE_BY_ID, data.Name, data.Description, time.Now().Unix(), data.ID)
}

func (r *role) DeleteRoleByID(ctx context.Context, id int64) error {
	return r.delete(ctx, QUERY_DELETE_ROLE_BY_ID, id)
}

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
	QUERY_FIND_BY_ID = `select id,uuid,username,password_hash,password_reset_token,email,nickname,avatar_url,status,
				create_at,update_at,delete_at from admin_user where id=?`
	QUERY_FIND_BY_UUID = `select id,uuid,username,password_hash,password_reset_token,email,nickname,avatar_url,status,
				create_at,update_at,delete_at from admin_user where uuid=?`
	QUERY_FIND_BY_USERNAME = `select id,uuid,username,password_hash,password_reset_token,email,nickname,avatar_url,status,
				create_at,update_at,delete_at from admin_user where username=?`
	QUERY_FIND_BY_EMAIL = `select id,uuid,username,password_hash,password_reset_token,email,nickname,avatar_url,status,
				create_at,update_at,delete_at from admin_user where email=?`
	QUERY_FIND_BY_MOBILE = `select id,uuid,username,password_hash,password_reset_token,email,nickname,avatar_url,status,
				create_at,update_at,delete_at from admin_user where mobile=?`
	QUERY_INSERT = `insert into admin_user ( uuid,username,password_hash,email,nickname,
				avatar_url,status,create_at ) values ( ?, ?, ?, ?, ?, ?, ?, ? )`
	QUERY_UPDATE_INFO_BY_UUID = `update admin_user set nickname=?,avatar_url=?,update_at=? where uuid=?`
	QUERY_UPDATE_PWD_BY_UUID  = `update admin_user set password_hash=?,update_at=? where uuid=?`
)

func NewAdminUserModel() model.AdminUserModel {
	return &adminUser{}
}

type adminUser struct {
}

func (this *adminUser) fetch(ctx context.Context, query string, args ...interface{}) ([]*dao.AdminUser, error) {
	rows, err := global.DB.QueryContext(ctx, query, args...)
	if err != nil {
		log.Error("admin_user fetch query error: ", err)
		return nil, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Error("admin_user fetch close error: ", err)
		}
	}()

	result := make([]*dao.AdminUser, 0)
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		row := new(dao.AdminUser)

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

func (this *adminUser) findUserRow(ctx context.Context, query string, args ...interface{}) (*dao.AdminUser, error) {
	list, err := this.fetch(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	if len(list) <= 0 {
		return nil, errors.New("未查询到数据")
	}
	return list[0], nil
}

func (this *adminUser) FindUserByID(ctx context.Context, id int64) (*dao.AdminUser, error) {
	return this.findUserRow(ctx, QUERY_FIND_BY_ID, id)
}

func (this *adminUser) FindUserByUUID(ctx context.Context, uuid string) (*dao.AdminUser, error) {
	return this.findUserRow(ctx, QUERY_FIND_BY_UUID, uuid)
}

func (this *adminUser) FindUserByUsername(ctx context.Context, username string) (*dao.AdminUser, error) {
	return this.findUserRow(ctx, QUERY_FIND_BY_USERNAME, username)
}

func (this *adminUser) FindUserByEmail(ctx context.Context, email string) (*dao.AdminUser, error) {
	return this.findUserRow(ctx, QUERY_FIND_BY_EMAIL, email)
}

func (this *adminUser) FindUserByMobile(ctx context.Context, mobile string) (*dao.AdminUser, error) {
	return this.findUserRow(ctx, QUERY_FIND_BY_MOBILE, mobile)
}

func (this *adminUser) InsertUser(ctx context.Context, user *dao.AdminUser) error {
	stmt, err := global.DB.PrepareContext(ctx, QUERY_INSERT)
	if err != nil {
		log.Error("admin_user insert prepare err: ", err)
		return err
	}

	res, err := stmt.ExecContext(ctx, user.UUID, user.UserName, user.PasswordHash, user.Email, user.Nickname, user.AvatarUrl, user.Status, time.Now().Unix())
	if err != nil {
		log.Error("admin_user insert error:", err)
		return err
	}

	user.ID, err = res.LastInsertId()
	if err != nil {
		log.Error("admin_user get insert id error:", err)
		return err
	}
	return nil
}

func (this *adminUser) updateUser(ctx context.Context, query string, args ...interface{}) error {
	stmt, err := global.DB.PrepareContext(ctx, query)
	if err != nil {
		log.Error("admin_user update error:", err)
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

func (this *adminUser) UpdateUserInfoByUUID(ctx context.Context, user *dao.AdminUser) error {
	return this.updateUser(ctx, QUERY_UPDATE_INFO_BY_UUID, user.Nickname, user.AvatarUrl, time.Now().Unix(), user.UUID)
}

func (a *adminUser) UpdateUserPwdByUUID(ctx context.Context, user *dao.AdminUser) error {
	return a.updateUser(ctx, QUERY_UPDATE_PWD_BY_UUID, user.PasswordHash, time.Now().Unix(), user.UUID)
}

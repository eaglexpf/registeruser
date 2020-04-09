package mysql

import (
	"context"
	"errors"
	"fmt"
	"registeruser/app/admin/entity"
	"registeruser/app/admin/model"
	"registeruser/conf/global"
	"registeruser/conf/log"
	"time"
)

const (
	QUERY_FIND_BY_ID = `select id,username,auth_key,password_hash,password_reset_token,email,nickname,avatar_url,status,
				create_at,update_at,delete_at from admin_user where id=?`
	QUERY_FIND_BY_UUID = `select id,username,auth_key,password_hash,password_reset_token,email,nickname,avatar_url,status,
				create_at,update_at,delete_at from admin_user where uuid=?`
	QUERY_FIND_BY_USERNAME = `select id,username,auth_key,password_hash,password_reset_token,email,nickname,avatar_url,status,
				create_at,update_at,delete_at from admin_user where username=?`
	QUERY_FIND_BY_EMAIL = `select id,username,auth_key,pass_hash,reset_pass_token,email,nickname,avatar_url,status,
				create_at,update_at,delete_at from admin_user where email=?`
	QUERY_FIND_BY_MOBILE = `select id,username,auth_key,pass_hash,reset_pass_token,email,nickname,avatar_url,status,
				create_at,update_at,delete_at from admin_user where mobile=?`
	QUERY_INSERT = `insert into admin_user ( username,auth_key,password_hash,password_reset_token,email,nickname,
				avatar_url,status,create_at ) values ( ?, ?, ?, ?, ?, ?, ?, ?, ? )`
	QUERY_UPDATE_INFO_BY_ID       = `update admin_user set nickname=?,avatar_url=?,update_at=? where id=?`
	QUERY_UPDATE_INFO_BY_UUID     = `update admin_user set nickname=?,avatar_url=?,update_at=? where uuid=?`
	QUERY_UPDATE_INFO_BY_USERNAME = `update admin_user set nickname=?,avatar_url=?,update_at=? where username=?`
	QUERY_UPDATE_INFO_BY_EMAIL    = `update admin_user set nickname=?,avatar_url=?,update_at=? where email=?`
	QUERY_UPDATE_INFO_BY_MOBILE   = `update admin_user set nickname=?,avatar_url=?,update_at=? where mobile=?`
)

func NewAdminUserModel() model.AdminUserModel {
	return &adminUser{}
}

type adminUser struct {
}

func (this *adminUser) fetch(ctx context.Context, query string, args ...interface{}) ([]*entity.AdminUser, error) {
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

	result := make([]*entity.AdminUser, 0)
	for rows.Next() {
		row := new(entity.AdminUser)
		err := rows.Scan(&row)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

func (this *adminUser) findUserRow(ctx context.Context, query string, args ...interface{}) (*entity.AdminUser, error) {
	list, err := this.fetch(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	if len(list) <= 0 {
		return nil, errors.New("未查询到数据")
	}
	return list[0], nil
}

func (this *adminUser) FindUserByID(ctx context.Context, id int64) (*entity.AdminUser, error) {
	return this.findUserRow(ctx, QUERY_FIND_BY_ID, id)
}

func (this *adminUser) FindUserByUUID(ctx context.Context, uuid string) (*entity.AdminUser, error) {
	return this.findUserRow(ctx, QUERY_FIND_BY_UUID, uuid)
}

func (this *adminUser) FindUserByUsername(ctx context.Context, username string) (*entity.AdminUser, error) {
	return this.findUserRow(ctx, QUERY_FIND_BY_USERNAME, username)
}

func (this *adminUser) FindUserByEmail(ctx context.Context, email string) (*entity.AdminUser, error) {
	return this.findUserRow(ctx, QUERY_FIND_BY_EMAIL, email)
}

func (this *adminUser) FindUserByMobile(ctx context.Context, mobile string) (*entity.AdminUser, error) {
	return this.findUserRow(ctx, QUERY_FIND_BY_MOBILE, mobile)
}

func (this *adminUser) InsertUser(ctx context.Context, user *entity.AdminUser) error {
	stmt, err := global.DB.PrepareContext(ctx, QUERY_INSERT)
	if err != nil {
		log.Error("admin_user insert prepare err: ", err)
		return err
	}

	res, err := stmt.Exec(ctx, user.UserName, user.Nickname)
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

func (this *adminUser) UpdateUserInfoByID(ctx context.Context, user entity.AdminUser) error {
	return this.updateUser(ctx, user.Nickname, user.AvatarUrl, time.Now().Unix(), user.ID)
}

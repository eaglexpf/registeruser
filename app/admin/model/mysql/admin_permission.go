package mysql

import (
	"context"
	"errors"
	"registeruser/app/admin/entity/dao"
	"registeruser/app/admin/model"
	"registeruser/conf/global"
	"registeruser/util/sql_util"
)

const (
	QUERY_PERMISSION_SEARCH_ALL               = `select * from admin_permission limit ?,?`
	QUERY_PERMISSION_SEARCH_INFO              = `select * from admin_permission where parent_id=? and parent_type=? and children_id=? and children_type=? limit ?,?`
	QUERY_PERMISSION_SEARCH_BY_PARENT         = `select * from admin_permission where parent_id=? and parent_type=? limit ?,?`
	QUERY_PERMISSION_SEARCH_BY_CHILDREN       = `select * from admin_permission where children_id=? and children_type=? limit ?,?`
	QUERY_PERMISSION_SEARCH_ALL_COUNT         = `select count(*) as count from admin_permission`
	QUERY_PERMISSION_SEARCH_INFO_COUNT        = `select count(*) as count from admin_permission where parent_id=? and parent_type=? and children_id=? and children_type=?`
	QUERY_PERMISSION_SEARCH_BY_PARENT_COUNT   = `select count(*) as count from admin_permission where parent_id=? and parent_type=?`
	QUERY_PERMISSION_SEARCH_BY_CHILDREN_COUNT = `select count(*) as count from admin_permission where children_id=? and children_type=?`
	QUERY_PERMISSION_REGISTER                 = `insert into admin_permission ( parent_id,parent_type,children_id,children_type ) values ( ?,?,?,? )`
	QUERY_PERMISSION_DELETE_INFO              = `delete from admin_permission where parent_id=? and parent_type=? and children_id=? and children_type=?`
	QUERY_PERMISSION_DELETE_BY_PARENT         = `delete from admin_permission where parent_id=? and parent_type=?`
	QUERY_PERMISSION_DELETE_BY_CHILDREN       = `delete from admin_permission where children_id=? and children_type=?`
)

func NewAdminPermissionModel() model.AdminPermissionModel {
	return &permissionModel{sql_util.NewSqlUtil(global.DB)}
}

type permissionModel struct {
	*sql_util.SqlUtil
}

func (p *permissionModel) check(data *dao.AdminPermission) error {
	if data.ParentType != "" {
		if data.ParentType != "user" && data.ParentType != "role" {
			return errors.New("parent_type类型只能为【user,role】")
		}
	}
	if data.ChildrenType != "" {
		if data.ChildrenType != "role" && data.ChildrenType != "api" {
			return errors.New("children_type类型只能为【role,api】")
		}
	}
	return nil
}

func (p *permissionModel) Search(ctx context.Context, search *dao.AdminPermission, offset, limit int64) (result []dao.AdminPermission, err error) {
	if err = p.check(search); err != nil {
		return
	}
	switch true {
	case search.ParentType != "" && search.ParentID > 0 && search.ChildrenType != "" && search.ChildrenID > 0:
		err = p.Fetch(ctx, QUERY_PERMISSION_SEARCH_INFO, &result, search.ParentID, search.ParentType, search.ChildrenID, search.ChildrenType, offset, limit)
	case search.ParentType != "" && search.ParentID > 0:
		err = p.Fetch(ctx, QUERY_PERMISSION_SEARCH_BY_PARENT, &result, search.ParentID, search.ParentType, offset, limit)
	case search.ChildrenType != "" && search.ChildrenID > 0:
		err = p.Fetch(ctx, QUERY_PERMISSION_SEARCH_BY_CHILDREN, &result, search.ChildrenID, search.ChildrenType, offset, limit)
	default:
		err = p.Fetch(ctx, QUERY_PERMISSION_SEARCH_ALL, &result, offset, limit)
	}
	return
}

func (p *permissionModel) SearchCount(ctx context.Context, search *dao.AdminPermission) int64 {
	var result map[string]interface{}
	var err error
	if err = p.check(search); err != nil {
		return 0
	}
	switch true {
	case search.ParentType != "" && search.ParentID > 0 && search.ChildrenType != "" && search.ChildrenID > 0:
		result, err = p.FetchMapRow(ctx, QUERY_PERMISSION_SEARCH_BY_PARENT_COUNT, search.ParentID, search.ParentType, search.ChildrenID, search.ChildrenType)
	case search.ParentType != "" && search.ParentID > 0:
		result, err = p.FetchMapRow(ctx, QUERY_PERMISSION_SEARCH_BY_PARENT_COUNT, search.ParentID, search.ParentType)
	case search.ChildrenType != "" && search.ChildrenID > 0:
		result, err = p.FetchMapRow(ctx, QUERY_PERMISSION_SEARCH_BY_CHILDREN_COUNT, search.ChildrenID, search.ChildrenType)
	default:
		result, err = p.FetchMapRow(ctx, QUERY_PERMISSION_SEARCH_ALL_COUNT)
	}
	if err != nil {
		return 0
	}
	if _, ok := result["count"]; !ok {
		return 0
	}
	return result["count"].(int64)
}

func (p *permissionModel) Register(ctx context.Context, data *dao.AdminPermission) error {
	if err := p.check(data); err != nil {
		return err
	}
	_, err := p.Insert(ctx, QUERY_PERMISSION_REGISTER, data.ParentID, data.ParentType, data.ChildrenID, data.ChildrenType)
	return err
}

func (p *permissionModel) DeleteInfo(ctx context.Context, data *dao.AdminPermission) (err error) {
	if err := p.check(data); err != nil {
		return err
	}
	switch true {
	case data.ParentType != "" && data.ParentID > 0 && data.ChildrenType != "" && data.ChildrenID > 0:
		_, err = p.Delete(ctx, QUERY_PERMISSION_DELETE_INFO, data.ParentID, data.ParentType, data.ChildrenID, data.ChildrenType)
	case data.ParentType != "" && data.ParentID > 0:
		_, err = p.Delete(ctx, QUERY_PERMISSION_DELETE_BY_PARENT, data.ParentID, data.ParentType)
	case data.ChildrenType != "" && data.ChildrenID > 0:
		_, err = p.Delete(ctx, QUERY_PERMISSION_DELETE_BY_CHILDREN, data.ChildrenID, data.ChildrenType)
	default:
		err = errors.New("无效的查询条件")
	}
	return
}

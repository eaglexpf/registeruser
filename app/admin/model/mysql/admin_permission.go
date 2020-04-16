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

	QUERY_PERMISSION_ROLE_CHILDREN = `select * from admin_permission where parent_id=? and parent_type="role"`
	QUERY_PERMISSION_USER_CHILDREN = `select * from admin_permission where parent_id=? and parent_type="user"`
)

func NewAdminPermissionModel() model.AdminPermissionModel {
	return &permissionModel{sql_util.NewSqlUtil(global.DB)}
}

type permissionModel struct {
	*sql_util.SqlUtil
}

// 检查parent_type和children_type是否正确
func (p *permissionModel) check(data *dao.AdminPermission) error {
	if data.ParentType != "" {
		if data.ParentType != "user" && data.ParentType != "role" {
			return errors.New("parent_type类型只能为【user,role】")
		}
	}
	if data.ChildrenType != "" {
		if data.ChildrenType != "role" && data.ChildrenType != "api" && data.ChildrenType != "service" {
			return errors.New("children_type类型只能为【role,api,service】")
		}
	}
	return nil
}

// 查询列表
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

// 按查询条件获取总条数
func (p *permissionModel) SearchCount(ctx context.Context, search *dao.AdminPermission) int64 {
	var result map[string]interface{}
	var err error
	if err = p.check(search); err != nil {
		return 0
	}
	switch true {
	case search.ParentType != "" && search.ParentID > 0 && search.ChildrenType != "" && search.ChildrenID > 0:
		result, err = p.FetchMapRow(ctx, QUERY_PERMISSION_SEARCH_INFO_COUNT, search.ParentID, search.ParentType, search.ChildrenID, search.ChildrenType)
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

// 注册
func (p *permissionModel) Register(ctx context.Context, data *dao.AdminPermission) error {
	if err := p.check(data); err != nil {
		return err
	}
	_, err := p.Insert(ctx, QUERY_PERMISSION_REGISTER, data.ParentID, data.ParentType, data.ChildrenID, data.ChildrenType)
	return err
}

// 删除
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

// 查询用户下的所有权限
func (p *permissionModel) FindPermissionListByUserID(ctx context.Context, id int64, apiIDS []int64) (result []int64) {
	data, err := p.FetchMap(ctx, QUERY_PERMISSION_USER_CHILDREN, id)
	if err != nil {
		return
	}
	if len(data) > 0 {
		for _, v := range data {
			children_type, ok := v["children_type"]
			if !ok {
				continue
			}
			children_type_str, ok := children_type.(string)
			if !ok {
				continue
			}
			children_id, ok := v["children_id"]
			if !ok {
				continue
			}
			children_id_int64, ok := children_id.(int64)
			if !ok {
				continue
			}
			//if children_type_str == "api" {
			//	apiIDS = append(apiIDS, children_id_int64)
			//	continue
			//}
			if children_type_str == "service" {
				apiIDS = append(apiIDS, children_id_int64)
				continue
			}
			if children_type_str == "role" {
				children_ids := p.FindPermissionListByRoleID(ctx, children_id_int64, apiIDS)
				apiIDS = append(apiIDS, children_ids...)
			}

		}

	}
	return apiIDS
}

// 查询角色下所有的权限
func (p *permissionModel) FindPermissionListByRoleID(ctx context.Context, id int64, apiIDS []int64) (result []int64) {
	data, err := p.FetchMap(ctx, QUERY_PERMISSION_ROLE_CHILDREN, id)
	if err != nil {
		return
	}
	if len(data) > 0 {
		for _, v := range data {
			children_type, ok := v["children_type"]
			if !ok {
				continue
			}
			children_type_str, ok := children_type.(string)
			if !ok {
				continue
			}
			children_id, ok := v["children_id"]
			if !ok {
				continue
			}
			children_id_int64, ok := children_id.(int64)
			if !ok {
				continue
			}
			//if children_type_str == "api" {
			//	apiIDS = append(apiIDS, children_id_int64)
			//	continue
			//}
			if children_type_str == "service" {
				apiIDS = append(apiIDS, children_id_int64)
				continue
			}
			if children_type_str == "role" {
				children_ids := p.FindPermissionListByRoleID(ctx, children_id_int64, apiIDS)
				apiIDS = append(apiIDS, children_ids...)
			}

		}

	}
	return apiIDS
}

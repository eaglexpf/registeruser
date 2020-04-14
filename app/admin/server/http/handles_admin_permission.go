package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"registeruser/app/admin/entity/request"
	"registeruser/app/admin/entity/response"
)

/**
 * @apiDefine api_group_104 4.权限管理
 */
func registerPermissionGroup(route *gin.RouterGroup) {
	permission := route.Group("/permission")
	permission.GET("/search", permissionSearch)
	permission.POST("/", permissionRegister)
	permission.PUT("/", permissionDelete)
}

/**
 * @api {get} /admin/permission/search 1.搜索权限
 * @apiDescription 按条件搜索权限列表
 * @apiGroup api_group_104
 * @apiVersion 1.0.0
 *
 * @apiHeader {string} token jwt验证token
 * @apiParam {int} [parent_id] 父权限id
 * @apiParam {string} [parent_type] 父权限类型
 * @apiParam {int} [children_id] 子权限id
 * @apiParam {string} [children_type] 子权限类型
 * @apiParam {int} [page=1] 第几页
 * @apiParam {int} [page_size=10] 每页多少条
 *
 * @apiSuccess {int} code 状态值
 * @apiSuccess {string} msg 状态描述
 * @apiSuccess {object} data 返回数据
 *
 **/
func permissionSearch(c *gin.Context) {
	var requestData request.RequestPermissionSearch
	if err := c.ShouldBind(&requestData); err != nil {
		c.JSON(http.StatusOK, response.ErrorParamValidateMsg(err.Error()))
		return
	}
	responseData, err := srv.PermissionSearch(c, &requestData)
	if err != nil {
		c.JSON(http.StatusOK, response.ErrorParamValidateMsg(err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.Success(responseData))
}

/**
 * @api {post} /admin/permission/ 2.添加权限
 * @apiDescription 注册一个新的权限
 * @apiGroup api_group_104
 * @apiVersion 1.0.0
 *
 * @apiHeader {string} token jwt验证token
 * @apiParam {int} parent_id 父权限id
 * @apiParam {string} parent_type 父权限类型
 * @apiParam {int} children_id 子权限id
 * @apiParam {string} children_type 子权限类型
 *
 * @apiSuccess {int} code 状态值
 * @apiSuccess {string} msg 状态描述
 * @apiSuccess {object} data 返回数据
 *
 **/
func permissionRegister(c *gin.Context) {
	var requestData request.RequestPermissionRegister
	if err := c.ShouldBind(&requestData); err != nil {
		c.JSON(http.StatusOK, response.ErrorParamValidateMsg(err.Error()))
		return
	}
	err := srv.PermissionRegister(c, &requestData)
	if err != nil {
		c.JSON(http.StatusOK, response.ErrorParamValidateMsg(err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.Success(gin.H{}))
}

/**
 * @api {put} /admin/permission/ 3.移除权限
 * @apiDescription 按条件删除权限，使用put方法，delete方法理论上只能通过url传参
 * @apiGroup api_group_104
 * @apiVersion 1.0.0
 *
 * @apiHeader {string} token jwt验证token
 * @apiParam {int} [parent_id] 父权限id
 * @apiParam {string} [parent_type] 父权限类型
 * @apiParam {int} [children_id] 子权限id
 * @apiParam {string} [children_type] 子权限类型
 *
 * @apiSuccess {int} code 状态值
 * @apiSuccess {string} msg 状态描述
 * @apiSuccess {object} data 返回数据
 *
 **/
func permissionDelete(c *gin.Context) {
	var requestData request.RequestPermissionDelete
	if err := c.ShouldBind(&requestData); err != nil {
		c.JSON(http.StatusOK, response.ErrorParamValidateMsg(err.Error()))
		return
	}
	err := srv.PermissionDelete(c, &requestData)
	if err != nil {
		c.JSON(http.StatusOK, response.ErrorParamValidateMsg(err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.Success(gin.H{}))
}

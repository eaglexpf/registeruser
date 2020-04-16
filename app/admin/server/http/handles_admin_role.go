package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"registeruser/app/admin/entity/request"
	"registeruser/app/admin/entity/response"
	"strconv"
)

/**
 * @apiDefine api_group_102 2.后台角色
 */

func registerRoleGroup(router *gin.RouterGroup) {
	role := router.Group("/role")
	role.GET("/", GinHandler(roleFindAll))
	role.POST("/", GinHandler(roleRegister))
	role.GET("/:id", GinHandler(roleFindByID))
	role.PUT("/:id", GinHandler(roleUpdateByID))
	role.DELETE("/:id", GinHandler(roleDeleteByID))
}

/**
 * @api {get} /admin/role/ 1.角色列表
 * @apiDescription 获取后台角色列表api
 * @apiGroup api_group_102
 * @apiVersion 1.0.0
 *
 * @apiHeader {string} token jwt验证token
 * @apiParam {int} [page=1] 第几页
 * @apiParam {int} [page_size=10] 每页多少条
 *
 * @apiSuccess {int} code 状态值
 * @apiSuccess {string} msg 状态描述
 * @apiSuccess {object} data 返回数据
 *
 **/
// 查询角色列表api
func roleFindAll(c *gin.Context) (err error) {
	var requestData request.RequestPage
	if err = c.ShouldBind(&requestData); err != nil {
		return
	}
	list, err := srv.FindRoleList(c, requestData.Page, requestData.PageSize)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, response.Success(list))
	return
}

/**
 * @api {get} /admin/role/:id 2.角色信息
 * @apiDescription 查询某个角色的借本信息
 * @apiGroup api_group_102
 * @apiVersion 1.0.0
 *
 * @apiHeader {string} token jwt验证token
 * @apiParam {int} id 角色id
 *
 * @apiSuccess {int} code 状态值
 * @apiSuccess {string} msg 状态描述
 * @apiSuccess {object} data 返回数据
 *
 **/
// 查询角色信息api
func roleFindByID(c *gin.Context) (err error) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		err = response.ERROR_PARAM_VALIDATE
		return
	}
	adminRole, err := srv.FindAdminRoleByID(c, id)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, response.Success(adminRole))
	return
}

/**
 * @api {post} /admin/role/ 3.新建角色
 * @apiDescription 注册一个新角色
 * @apiGroup api_group_102
 * @apiVersion 1.0.0
 *
 * @apiHeader {string} token jwt验证token
 * @apiParam {string} name 角色名称
 * @apiParam {string} description 角色描述
 *
 * @apiSuccess {int} code 状态值
 * @apiSuccess {string} msg 状态描述
 * @apiSuccess {object} data 返回数据
 *
 **/
// 注册一个新角色
func roleRegister(c *gin.Context) (err error) {
	request_admin_role := new(request.RequestAdminRoleRegister)
	if err = c.ShouldBind(request_admin_role); err != nil {
		return
	}
	rsponse_admin_role, err := srv.RegisterAdminRole(c, request_admin_role)
	if err != nil {
		return
	}
	c.JSON(200, response.Success(rsponse_admin_role))
	return
}

/**
 * @api {put} /admin/role/:id 4.修改角色
 * @apiDescription 注册一个新角色
 * @apiGroup api_group_102
 * @apiVersion 1.0.0
 *
 * @apiHeader {string} token jwt验证token
 * @apiParam {int} id 角色id
 * @apiParam {string} name 角色名称
 * @apiParam {string} description 角色描述
 *
 * @apiSuccess {int} code 状态值
 * @apiSuccess {string} msg 状态描述
 * @apiSuccess {object} data 返回数据
 *
 **/
// 修改一个角色
func roleUpdateByID(c *gin.Context) (err error) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		err = response.ERROR_PARAM_VALIDATE
		return
	}
	request_admin_role_update := new(request.RequestAdminRoleUpdate)
	if err = c.ShouldBind(request_admin_role_update); err != nil {
		return
	}
	request_admin_role_update.ID = id
	response_admin_role, err := srv.UpdateRoleByID(c, request_admin_role_update)
	if err != nil {
		return
	}
	c.JSON(200, response.Success(response_admin_role))
	return
}

/**
 * @api {delete} /admin/role/:id 5.删除角色
 * @apiDescription 删除一个角色
 * @apiGroup api_group_102
 * @apiVersion 1.0.0
 *
 * @apiHeader {string} token jwt验证token
 * @apiParam {int} id 角色id
 *
 * @apiSuccess {int} code 状态值
 * @apiSuccess {string} msg 状态描述
 * @apiSuccess {object} data 返回数据
 *
 **/
// 删除一个角色
func roleDeleteByID(c *gin.Context) (err error) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		err = response.ERROR_PARAM_VALIDATE
		return
	}
	err = srv.DeleteAdminRoleByID(c, id)
	if err != nil {
		return
	}
	c.JSON(200, response.Success(nil))
	return
}

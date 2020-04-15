package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"registeruser/app/admin/entity/request"
	"registeruser/app/admin/entity/response"
	"strconv"
)

/**
 * @apiDefine api_group_103 3.服务管理
 */

func registerServiceGroup(router *gin.RouterGroup) {
	service := router.Group("/service")
	service.GET("/row/:id", adminServiceFindByID)
	service.GET("/search", adminServiceSearch)
	service.POST("/", adminServiceRegister)
	service.PUT("/row/:id", adminServiceUpdate)
	service.DELETE("/id/:id", adminServiceDelByID)
	service.DELETE("/name/:name", adminServiceDelByName)
}

/**
 * @api {get} /admin/service/row/:id 1.服务信息
 * @apiDescription 按id查询服务信息
 * @apiGroup api_group_103
 * @apiVersion 1.0.0
 *
 * @apiHeader {string} token jwt验证token
 * @apiParam {int} id 服务id
 *
 * @apiSuccess {int} code 状态值
 * @apiSuccess {string} msg 状态描述
 * @apiSuccess {object} data 返回数据
 *
 **/
func adminServiceFindByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, response.ErrorParamValidateMsg(err.Error()))
		return
	}
	data, err := srv.ServiceFindByID(c, id)
	if err != nil {
		c.JSON(http.StatusOK, response.ErrorParamValidateMsg(err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.Success(data))
	return
}

/**
 * @api {get} /admin/service/search 2.服务列表
 * @apiDescription 搜索服务
 * @apiGroup api_group_103
 * @apiVersion 1.0.0
 *
 * @apiHeader {string} token jwt验证token
 * @apiParam {string} [name] 服务名称
 * @apiParam {string} [alias] 服务别名
 * @apiParam {int} [status] 服务状态[1：有效；2：停用]
 * @apiParam {int} [page=1] 第几页
 * @apiParam {int} [page_size=10] 每页多少条
 *
 * @apiSuccess {int} code 状态值
 * @apiSuccess {string} msg 状态描述
 * @apiSuccess {object} data 返回数据
 *
 **/
func adminServiceSearch(c *gin.Context) {
	var search request.RequestServiceSearch
	if err := c.ShouldBind(&search); err != nil {
		c.JSON(http.StatusOK, response.ErrorParamValidateMsg(err.Error()))
		return
	}
	data, err := srv.ServiceSearch(c, &search)
	if err != nil {
		c.JSON(http.StatusOK, response.ErrorParamValidateMsg(err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.Success(data))
	return
}

/**
 * @api {post} /admin/service/ 3.注册服务
 * @apiDescription 注册服务
 * @apiGroup api_group_103
 * @apiVersion 1.0.0
 *
 * @apiHeader {string} token jwt验证token
 * @apiParam {string} name 服务名称
 * @apiParam {string} alias 服务别名
 * @apiParam {string} description 服务别名
 * @apiParam {int} status 服务状态[1：有效；2：停用]
 * @apiParam {int} expire_at 过期时间[时间戳格式，0为永不过期]
 *
 * @apiSuccess {int} code 状态值
 * @apiSuccess {string} msg 状态描述
 * @apiSuccess {object} data 返回数据
 *
 **/
func adminServiceRegister(c *gin.Context) {
	var data request.RequestServiceRegister
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusOK, response.ErrorParamValidateMsg(err.Error()))
		return
	}
	err := srv.ServiceRegister(c, &data)
	if err != nil {
		c.JSON(http.StatusOK, response.ErrorParamValidateMsg(err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.Success(struct{}{}))
	return
}

/**
 * @api {put} /admin/service/row/:id 4.修改服务
 * @apiDescription 注册服务
 * @apiGroup api_group_103
 * @apiVersion 1.0.0
 *
 * @apiHeader {string} token jwt验证token
 * @apiParam {int} id 服务id
 * @apiParam {string} name 服务名称
 * @apiParam {string} alias 服务别名
 * @apiParam {string} description 服务别名
 * @apiParam {int} status 服务状态[1：有效；2：停用]
 * @apiParam {int} expire_at 过期时间[时间戳格式，0为永不过期]
 *
 * @apiSuccess {int} code 状态值
 * @apiSuccess {string} msg 状态描述
 * @apiSuccess {object} data 返回数据
 *
 **/
func adminServiceUpdate(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, response.ErrorParamValidateMsg(err.Error()))
		return
	}
	var data request.RequestServiceUpdate
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusOK, response.ErrorParamValidateMsg(err.Error()))
		return
	}
	err = srv.ServiceUpdateByID(c, &data, id)
	if err != nil {
		c.JSON(http.StatusOK, response.ErrorParamValidateMsg(err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.Success(struct{}{}))
	return
}

/**
 * @api {delete} /admin/service/id/:id 5.删除服务
 * @apiDescription 根据id删除服务
 * @apiGroup api_group_103
 * @apiVersion 1.0.0
 *
 * @apiHeader {string} token jwt验证token
 * @apiParam {int} id 服务id
 *
 * @apiSuccess {int} code 状态值
 * @apiSuccess {string} msg 状态描述
 * @apiSuccess {object} data 返回数据
 *
 **/
func adminServiceDelByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, response.ErrorParamValidateMsg(err.Error()))
		return
	}
	err = srv.ServiceDeleteByID(c, id)
	if err != nil {
		c.JSON(http.StatusOK, response.ErrorParamValidateMsg(err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.Success(struct{}{}))
	return
}

/**
 * @api {delete} /admin/service/name/:name 6.删除服务-name
 * @apiDescription 根据name删除服务
 * @apiGroup api_group_103
 * @apiVersion 1.0.0
 *
 * @apiHeader {string} token jwt验证token
 * @apiParam {string} name 服务名称
 *
 * @apiSuccess {int} code 状态值
 * @apiSuccess {string} msg 状态描述
 * @apiSuccess {object} data 返回数据
 *
 **/
func adminServiceDelByName(c *gin.Context) {
	name := c.Param("name")
	err := srv.ServiceDeleteByName(c, name)
	if err != nil {
		c.JSON(http.StatusOK, response.ErrorParamValidateMsg(err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.Success(struct{}{}))
	return
}

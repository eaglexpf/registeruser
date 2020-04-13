package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"registeruser/app/admin/entity/request"
	"registeruser/app/admin/entity/response"
	"strconv"
)

/**
 * @apiDefine api_group_103 2.API管理
 */

func registerApiGroup(router *gin.RouterGroup) {
	api := router.Group("api")
	api.GET("/", apiFindAll)
	api.POST("/", apiRegister)
	api.GET("/id/:id", apiFindByID)
	api.PUT("/id/:id", apiUpdateByID)
	api.DELETE("/id/:id", apiDeleteByID)
	api.GET("/search", apiFindBySearch)
}

/**
 * @api {get} /admin/api/ 1.API列表
 * @apiDescription 获取API列表
 * @apiGroup api_group_103
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
// 查询api列表
func apiFindAll(c *gin.Context) {
	page, err := strconv.ParseInt(c.DefaultQuery("page", "1"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, response.ErrorParamValidateData(err.Error()))
		return
	}
	page_size, err := strconv.ParseInt(c.DefaultQuery("page_size", "10"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, response.ErrorParamValidateData(err.Error()))
		return
	}
	list, err := srv.ApiFindAll(c, page, page_size)
	if err != nil {
		c.JSON(http.StatusOK, response.ErrorParamValidateData(err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.Success(list))
	return
}

/**
 * @api {get} /admin/api/id/:id 2.根据id查询
 * @apiDescription 根据id查询api
 * @apiGroup api_group_103
 * @apiVersion 1.0.0
 *
 * @apiHeader {string} token jwt验证token
 * @apiParam {int} id api的id
 *
 * @apiSuccess {int} code 状态值
 * @apiSuccess {string} msg 状态描述
 * @apiSuccess {object} data 返回数据
 *
 **/
// 根据id查询api
func apiFindByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, response.ErrorParamValidateMsg(err.Error()))
		return
	}
	data, err := srv.ApiFindByID(c, id)
	if err != nil {
		c.JSON(http.StatusOK, response.ErrorParamValidateMsg(err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.Success(data))
	return
}

/**
 * @api {get} /admin/api/search 3.搜索
 * @apiDescription 搜索api；path模糊搜索
 * @apiGroup api_group_103
 * @apiVersion 1.0.0
 *
 * @apiHeader {string} token jwt验证token
 * @apiParam {string} method 请求方式
 * @apiParam {string} path api路径
 * @apiParam {int} [page=1] 第几页
 * @apiParam {int} [page_size=10] 每页多少条
 *
 * @apiSuccess {int} code 状态值
 * @apiSuccess {string} msg 状态描述
 * @apiSuccess {object} data 返回数据
 *
 **/
// 搜索
func apiFindBySearch(c *gin.Context) {
	page, err := strconv.ParseInt(c.DefaultQuery("page", "1"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, response.ErrorParamValidateData(err.Error()))
		return
	}
	page_size, err := strconv.ParseInt(c.DefaultQuery("page_size", "10"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, response.ErrorParamValidateData(err.Error()))
		return
	}
	method := c.Query("method")
	path := c.Query("path")
	list, err := srv.ApiSearch(c, method, path, page, page_size)
	if err != nil {
		c.JSON(http.StatusOK, response.ErrorParamValidateMsg(err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.Success(list))
	return
}

/**
 * @api {post} /admin/api/ 4.注册新api
 * @apiDescription 创建一个新的api
 * @apiGroup api_group_103
 * @apiVersion 1.0.0
 *
 * @apiHeader {string} token jwt验证token
 * @apiParam {string} method 请求方式
 * @apiParam {string} path api路径
 * @apiParam {string} description api描述
 *
 * @apiSuccess {int} code 状态值
 * @apiSuccess {string} msg 状态描述
 * @apiSuccess {object} data 返回数据
 *
 **/
func apiRegister(c *gin.Context) {
	request_data := new(request.RequestApiRegister)
	if err := c.ShouldBind(request_data); err != nil {
		c.JSON(http.StatusOK, response.ErrorParamValidateMsg(err.Error()))
		return
	}
	err := srv.ApiRegister(c, request_data)
	if err != nil {
		c.JSON(http.StatusOK, response.ErrorParamValidateMsg(err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.Success(nil))
	return
}

/**
 * @api {put} /admin/api/id/:id 5.修改api
 * @apiDescription 创建一个新的api
 * @apiGroup api_group_103
 * @apiVersion 1.0.0
 *
 * @apiHeader {string} token jwt验证token
 * @apiParam {int} id api的id
 * @apiParam {string} method 请求方式
 * @apiParam {string} path api路径
 * @apiParam {string} description api描述
 *
 * @apiSuccess {int} code 状态值
 * @apiSuccess {string} msg 状态描述
 * @apiSuccess {object} data 返回数据
 *
 **/
func apiUpdateByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, response.ErrorParamValidateMsg("id不能为空"))
		return
	}
	request_data := new(request.RequestApiUpdate)
	if err := c.ShouldBind(request_data); err != nil {
		c.JSON(http.StatusOK, response.ErrorParamValidateMsg(err.Error()))
		return
	}
	err = srv.ApiUpdateByID(c, request_data, id)
	if err != nil {
		c.JSON(http.StatusOK, response.ErrorParamValidateMsg(err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.Success(nil))
	return
}

/**
 * @api {delete} /admin/api/id/:id 6.删除api
 * @apiDescription 根据id删除api
 * @apiGroup api_group_103
 * @apiVersion 1.0.0
 *
 * @apiHeader {string} token jwt验证token
 * @apiParam {int} id api的id
 *
 * @apiSuccess {int} code 状态值
 * @apiSuccess {string} msg 状态描述
 * @apiSuccess {object} data 返回数据
 *
 **/
func apiDeleteByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, response.ErrorParamValidateMsg(err.Error()))
		return
	}
	err = srv.ApiDeleteByID(c, id)
	if err != nil {
		c.JSON(http.StatusOK, response.ErrorParamValidateMsg(err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.Success(nil))
	return
}

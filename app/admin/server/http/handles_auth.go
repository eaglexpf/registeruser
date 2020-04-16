package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"registeruser/app/admin/entity/response"
)

func registerAuthGroup(router *gin.RouterGroup) {
	auth := router.Group("/auth")
	auth.GET("/:name", GinHandler(authCheck))
}

/**
 * @apiDefine api_group_105 5.权限校验
 */

/**
 * @api {get} /admin/auth/:name 1.权限校验
 * @apiDescription 权限校验
 * @apiGroup api_group_105
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
func authCheck(c *gin.Context) (err error) {
	name := c.Param("name")
	err = srv.AuthCheck(c, name)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	c.JSON(http.StatusOK, response.Success(struct{}{}))
	return
}

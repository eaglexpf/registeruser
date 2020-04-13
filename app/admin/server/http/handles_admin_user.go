package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"registeruser/app/admin/entity/dao"
	"registeruser/app/admin/entity/request"
	"registeruser/app/admin/entity/response"
)

func adminUserList(c *gin.Context) {
	c.JSON(200, response.Success(nil))
}

/**
 * @apiDefine api_group_101 1.后台用户
 */

/**
 * @api {post} /admin/user/register 1.注册后台用户api
 * @apiDescription 新建一个新的后台用户
 * @apiGroup api_group_101
 * @apiVersion 1.0.0
 *
 * @apiHeader {string} token jwt验证token
 *
 * @apiParam {string} username 账号
 * @apiParam {string} password 密码
 * @apiParam {string} repeat_pwd 再输入一次密码
 * @apiParam {string} email 邮箱
 * @apiParam {string} nickname 昵称
 * @apiParam {string} avatar_url 头像
 *
 * @apiSuccess {int} code 状态值
 * @apiSuccess {string} msg 状态描述
 * @apiSuccess {object} data 返回数据
 *
 **/

// 注册后台用户api
func adminUserRegister(c *gin.Context) {
	adminUser := new(request.RequestRegisterAdminUser)
	if err := c.ShouldBind(adminUser); err != nil {
		c.JSON(http.StatusOK, response.ErrorParamValidateData(err.Error()))
		return
	}
	_, err := srv.RegisterAdminUser(c, adminUser)
	if err != nil {
		c.JSON(http.StatusOK, response.Error(400, err.Error()))
		return
	}

	c.JSON(200, response.Success(nil))
}

/**
 * @api {post} /admin/user/login 2.后台用户登录
 * @apiDescription 后台用户登录
 * @apiGroup api_group_101
 * @apiVersion 1.0.0
 *
 * @apiParam {string} username 账号
 * @apiParam {string} password 密码
 *
 * @apiSuccess {int} code 状态值
 * @apiSuccess {string} msg 状态描述
 * @apiSuccess {object} data 返回数据
 *
 **/

// 后台用户登录api
func adminUserLogin(c *gin.Context) {
	adminUser := new(request.RequestAdminUserLogin)
	if err := c.ShouldBind(adminUser); err != nil {
		c.JSON(http.StatusOK, response.ErrorParamValidateData(err.Error()))
	}
	token, err := srv.LoginForAdminUser(c, adminUser)
	if err != nil {
		c.JSON(http.StatusOK, response.ErrorRegisterAdminUser())
	}
	c.JSON(http.StatusOK, response.Success(token))
}

/**
 * @api {get} /admin/user/refresh 3.后台用户刷新token
 * @apiDescription 后台用户刷新token
 * @apiGroup api_group_101
 * @apiVersion 1.0.0
 *
 * @apiHeader {string} token jwt验证token
 *
 * @apiSuccess {int} code 状态值
 * @apiSuccess {string} msg 状态描述
 * @apiSuccess {object} data 返回数据
 *
 **/

// 后台用户刷新token
func adminUserRefreshToken(c *gin.Context) {
	user, ok := c.Get("adminUser")
	if !ok {
		c.JSON(http.StatusForbidden, response.Error(http.StatusForbidden, "adminUser获取失败"))
		return
	}
	adminUser, ok := user.(*dao.AdminUser)
	if !ok {
		c.JSON(http.StatusForbidden, response.Error(http.StatusForbidden, "adminUser类型获取失败"))
		return
	}

	token, err := srv.RefreshTokenByAdminUser(c, adminUser)
	if err != nil {
		c.JSON(http.StatusOK, response.ErrorRegisterAdminUser())
	}
	c.JSON(http.StatusOK, response.Success(token))
}

/**
 * @api {post} /admin/user/update-info 4.修改后台用户基本信息
 * @apiDescription 修改后台用户昵称，头像等基本信息
 * @apiGroup api_group_101
 * @apiVersion 1.0.0
 *
 * @apiHeader {string} token jwt验证token
 * @apiParam {string} nickname 昵称
 * @apiParam {string} avatar_url 头像
 *
 * @apiSuccess {int} code 状态值
 * @apiSuccess {string} msg 状态描述
 * @apiSuccess {object} data 返回数据
 *
 **/

// 修改后台用户基本信息
func adminUserUpdateInfo(c *gin.Context) {
	user, ok := c.Get("adminUser")
	if !ok {
		c.JSON(http.StatusForbidden, response.Error(http.StatusForbidden, "adminUser获取失败"))
		return
	}
	adminUser, ok := user.(*dao.AdminUser)
	if !ok {
		c.JSON(http.StatusForbidden, response.Error(http.StatusForbidden, "adminUser类型获取失败"))
		return
	}

	var updateData request.RequestAdminUserUpdateInfo
	if err := c.ShouldBind(&updateData); err != nil {
		c.JSON(http.StatusOK, response.ErrorParamValidateData(err.Error()))
		return
	}
	if updateData.UUID == "" {
		updateData.UUID = adminUser.UUID
	}

	adminUser, err := srv.UpdateAdminUserInfo(c, &updateData)
	if err != nil {
		c.JSON(http.StatusOK, response.ErrorParamValidateData(err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.Success(adminUser))
}

/**
 * @api {post} /admin/user/update-pwd 5.修改后台用户密码
 * @apiDescription 修改后台用户密码
 * @apiGroup api_group_101
 * @apiVersion 1.0.0
 *
 * @apiHeader {string} token jwt验证token
 * @apiParam {string} password 新密码
 * @apiParam {string} old_pwd 旧密码
 *
 * @apiSuccess {int} code 状态值
 * @apiSuccess {string} msg 状态描述
 * @apiSuccess {object} data 返回数据
 *
 **/

// 修改后台用户密码
func adminUserUpdatePwd(c *gin.Context) {
	user, ok := c.Get("adminUser")
	if !ok {
		c.JSON(http.StatusForbidden, response.Error(http.StatusForbidden, "adminUser获取失败"))
		return
	}
	adminUser, ok := user.(*dao.AdminUser)
	if !ok {
		c.JSON(http.StatusForbidden, response.Error(http.StatusForbidden, "adminUser类型获取失败"))
		return
	}

	var updateData request.RequestAdminUserResetPwd
	if err := c.ShouldBind(&updateData); err != nil {
		c.JSON(http.StatusOK, response.ErrorParamValidateData(err.Error()))
	}
	if updateData.UUID == "" {
		updateData.UUID = adminUser.UUID
	}

	adminUser, err := srv.ResetPwdForAdminUser(c, &updateData)
	if err != nil {
		c.JSON(http.StatusOK, response.ErrorParamValidateData(err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.Success(adminUser))
}

func adminUserResetEmail(c *gin.Context) {

}

func adminUserResetMobile(c *gin.Context) {

}

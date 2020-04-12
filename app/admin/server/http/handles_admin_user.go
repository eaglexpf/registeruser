package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"registeruser/app/admin/entity"
	"registeruser/app/admin/entity/request"
	"registeruser/app/admin/entity/response"
)

func adminUserList(c *gin.Context) {
	c.JSON(200, response.Success(nil))
}

func adminUserRegister(c *gin.Context) {
	adminUser := new(request.RequestRegisterAdminUser)
	if err := c.ShouldBind(adminUser); err != nil {
		c.JSON(http.StatusOK, response.ErrorParamValidateData(err.Error()))
		return
	}
	_, err := srv.AdminUserRegister(c, adminUser)
	if err != nil {
		c.JSON(http.StatusOK, response.Error(400, err.Error()))
		return
	}

	c.JSON(200, response.Success(nil))
}

func adminUserLogin(c *gin.Context) {
	adminUser := new(request.RequestAdminUserLogin)
	if err := c.ShouldBind(adminUser); err != nil {
		c.JSON(http.StatusOK, response.ErrorParamValidateData(err.Error()))
	}
	token, err := srv.AdminUserLogin(c, adminUser)
	if err != nil {
		c.JSON(http.StatusOK, response.ErrorRegisterAdminUser())
	}
	c.JSON(http.StatusOK, response.Success(token))
}

func adminUserRefreshToken(c *gin.Context) {
	user, ok := c.Get("adminUser")
	if !ok {
		c.JSON(http.StatusForbidden, response.Error(http.StatusForbidden, "adminUser获取失败"))
		return
	}
	adminUser, ok := user.(*entity.AdminUser)
	if !ok {
		c.JSON(http.StatusForbidden, response.Error(http.StatusForbidden, "adminUser类型获取失败"))
		return
	}

	token, err := srv.AdminUserRefreshToken(c, adminUser)
	if err != nil {
		c.JSON(http.StatusOK, response.ErrorRegisterAdminUser())
	}
	c.JSON(http.StatusOK, response.Success(token))
}

func adminUserUpdateInfo(c *gin.Context) {
	user, ok := c.Get("adminUser")
	if !ok {
		c.JSON(http.StatusForbidden, response.Error(http.StatusForbidden, "adminUser获取失败"))
		return
	}
	adminUser, ok := user.(*entity.AdminUser)
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

	adminUser, err := srv.AdminUserUpdateInfo(c, &updateData)
	if err != nil {
		c.JSON(http.StatusOK, response.ErrorParamValidateData(err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.Success(adminUser))
}

func adminUserUpdatePwd(c *gin.Context) {
	user, ok := c.Get("adminUser")
	if !ok {
		c.JSON(http.StatusForbidden, response.Error(http.StatusForbidden, "adminUser获取失败"))
		return
	}
	adminUser, ok := user.(*entity.AdminUser)
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

	adminUser, err := srv.AdminUserResetPwd(c, &updateData)
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

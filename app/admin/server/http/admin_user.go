package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"registeruser/app/admin/entity/request"
	"registeruser/app/admin/entity/response"
)

func adminUserList(c *gin.Context) {
	c.JSON(200, response.Success(nil))
}

func registerAdminUser(c *gin.Context) {
	adminUser := new(request.RequestRegisterAdminUser)
	if err := c.ShouldBind(adminUser); err != nil {
		c.JSON(http.StatusOK, response.ErrorParamValidateData(err.Error()))
		return
	}
	_, err := srv.RegisterUser(c, adminUser)
	if err != nil {
		c.JSON(http.StatusOK, response.Error(400, err.Error()))
		return
	}

	c.JSON(200, response.Success(nil))
}

func adminUserLogin(c *gin.Context) {

}

func adminUserRefreshToken(c *gin.Context) {

}

func adminUserLogout(c *gin.Context) {

}

func adminUserGetMobileCode(c *gin.Context) {

}

func adminUserRegisterMobile(c *gin.Context) {

}

func adminUserResetEmail(c *gin.Context) {

}

func adminUserResetMobile(c *gin.Context) {

}

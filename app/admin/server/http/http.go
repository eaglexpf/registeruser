package http

import (
	"github.com/gin-gonic/gin"
	"registeruser/app/admin/service"
)

var srv *service.Service

func init() {
	srv = service.NewService()
}

func Register(r *gin.Engine) *gin.Engine {
	router := r.Group("/admin/")
	user := router.Group("user")
	user.GET("/", adminUserList)
	user.POST("register", registerAdminUser)
	//srv := service.NewService()
	//router.Any("/", func(c *gin.Context) {
	//	fmt.Println("admin get")
	//	//jwt := NewJWT()
	//	//token, err := util.NewJWT().CreateToken(&global.JwtClaims{})
	//	var user request.RequestRegisterAdminUser
	//	if err := c.ShouldBind(&user); err != nil {
	//		c.JSON(200, response.ErrorParamValidateMsg(fmt.Sprint(err)))
	//		return
	//	}
	//	token, err := srv.FindUserByUsername(c, c.DefaultQuery("username", "admin"))
	//	if err == nil {
	//		// 账号已存在
	//		c.JSON(200, response.ErrorFindAdminUser())
	//		return
	//	}
	//	c.JSON(200, response.Success(token))
	//})
	return r

}

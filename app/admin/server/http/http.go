// 子项目的http服务包
package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"registeruser/app/admin/entity/response"
	"registeruser/app/admin/service"
	"registeruser/util/gin_middleware"
)

//var srv service.ServiceInter
var srv *service.Service

func init() {
	srv = service.NewService()
}

func GinHandler(handler func(c *gin.Context) error) func(c *gin.Context) {
	return func(c *gin.Context) {
		if err := handler(c); err != nil {
			result := response.GetErrorResponse(err)
			c.AbortWithStatusJSON(http.StatusOK, result)
		}
	}
}

// 注册gin路由
func Register(r *gin.Engine) {
	router := r.Group("/admin")
	user := router.Group("/user")
	user.POST("/login", GinHandler(adminUserLogin))
	user.Use(gin_middleware.JWTMiddleware()).Use(middlewareAdminUser())
	{
		user.POST("/register", GinHandler(adminUserRegister))
		user.GET("/refresh", GinHandler(adminUserRefreshToken))
		user.POST("/update-info", GinHandler(adminUserUpdateInfo))
		user.POST("/update-pwd", GinHandler(adminUserUpdatePwd))
	}
	router.Use(gin_middleware.JWTMiddleware()).Use(middlewareAdminUser())
	{

		registerRoleGroup(router)
		registerApiGroup(router)
		registerServiceGroup(router)
		registerPermissionGroup(router)
		registerAuthGroup(router)
	}
}

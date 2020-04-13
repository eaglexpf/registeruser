// 子项目的http服务包
package http

import (
	"github.com/gin-gonic/gin"
	"registeruser/app/admin/service"
	"registeruser/util/gin_middleware"
)

var srv *service.Service

func init() {
	srv = service.NewService()
}

// 注册gin路由
func Register(r *gin.Engine) {
	router := r.Group("/admin/")
	user := router.Group("user/")
	user.POST("register", adminUserRegister)
	user.POST("login", adminUserLogin)
	user.Use(gin_middleware.JWTMiddleware()).Use(middlewareAdminUser())
	{
		user.GET("refresh", adminUserRefreshToken)
		user.POST("update-info", adminUserUpdateInfo)
		user.POST("update-pwd", adminUserUpdatePwd)

		registerRoleGroup(router)
		registerApiGroup(router)
	}
}

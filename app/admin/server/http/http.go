package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"registeruser/app/admin/entity"
	"registeruser/app/admin/service"
)

func Register(r *gin.Engine) *gin.Engine {
	srv := service.NewService()
	router := r.Group("/admin")
	router.GET("/", func(c *gin.Context) {
		//jwt := NewJWT()
		//c.ShouldBindWith()
		//token, err := util.NewJWT().CreateToken(&global.JwtClaims{})
		var user entity.RequestAdminInsert
		if err := c.ShouldBind(&user); err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		token, err := srv.FindUserByUsername(c, c.DefaultQuery("username", "admin"))
		if err != nil {
			c.JSON(400, gin.H{
				"err": fmt.Sprint(err),
			})
			return
		}
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  token,
			"err":  err,
		})
	})
	return r

}

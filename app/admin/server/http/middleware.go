package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"registeruser/conf/global"
	"registeruser/entity/response"
)

func middlewareAdminUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, ok := c.Get("claims")
		if !ok {
			c.JSON(http.StatusForbidden, response.ErrorForBidden())
			c.Abort()
			return
		}
		uuid, ok := claims.(*global.JwtClaims)
		if !ok {
			c.JSON(http.StatusForbidden, response.ErrorForBidden())
			c.Abort()
			return
		}
		adminUser, err := srv.AdminUserFindByUUID(c, uuid.UUID)
		if err != nil {
			c.JSON(http.StatusForbidden, response.ErrorForBidden())
			c.Abort()
			return
		}
		c.Set("adminUser", adminUser)
		c.Next()
	}
}

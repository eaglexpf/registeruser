package ginMiddleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"registeruser/entity/response"
	"registeruser/util"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			c.JSON(http.StatusForbidden, response.ErrorForBidden())
			c.Abort()
			return
		}
		// parseToken 解析token包含的信息
		jwt := util.NewJWT()
		claims, err := jwt.ParseToken(token)
		if err != nil {
			if err == util.TokenExpired {
				c.JSON(http.StatusForbidden, response.ErrorForBidden())
				c.Abort()
				return
			}
			c.JSON(http.StatusForbidden, response.ErrorForBidden())
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}

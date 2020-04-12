package ginMiddleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"registeruser/entity/response"
	jwt2 "registeruser/util/jwt"
)

// gin的jwt验证中间件
func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			c.JSON(http.StatusForbidden, response.ErrorForBidden())
			c.Abort()
			return
		}
		// parseToken 解析token包含的信息
		jwt := jwt2.NewJWT()
		claims, err := jwt.ParseToken(token)
		if err != nil {
			if err == jwt2.TokenExpired {
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

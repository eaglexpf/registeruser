package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"registeruser/util"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			c.JSON(http.StatusForbidden, gin.H{
				"code": http.StatusForbidden,
				"msg":  "未授权的访问",
				"data": gin.H{},
			})
			c.Abort()
			return
		}
		// parseToken 解析token包含的信息
		jwt := util.NewJWT()
		claims, err := jwt.ParseToken(token)
		if err != nil {
			if err == util.TokenExpired {
				c.JSON(http.StatusForbidden, gin.H{
					"code": http.StatusForbidden,
					"msg":  "授权已过期",
					"data": gin.H{},
				})
				c.Abort()
				return
			}
			c.JSON(http.StatusForbidden, gin.H{
				"code": http.StatusForbidden,
				"msg":  "请重新授权",
				"data": gin.H{},
			})
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}

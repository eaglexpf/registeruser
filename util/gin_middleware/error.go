package gin_middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"registeruser/conf/log"
	"registeruser/entity/response"
)

// gin处理panic中间件
func RecoverMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 记录一个错误的日志
				log.Errorf("捕获到一个致命错误：", err)
				c.JSON(http.StatusServiceUnavailable, response.ErrorServiceUnavailable())
				return
			}
		}()
		c.Next()
	}
}

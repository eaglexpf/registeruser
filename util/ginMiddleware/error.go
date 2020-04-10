package ginMiddleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"registeruser/entity/response"
)

func RecoverMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 记录一个错误的日志
				c.JSON(http.StatusServiceUnavailable, response.ErrorServiceUnavailable(fmt.Sprintf("捕获到一个致命错误：%v", err)))
				return
			}
		}()
		c.Next()
	}
}

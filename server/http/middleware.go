package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"registeruser/conf/log"
	"registeruser/util"
	"time"
)

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

func jwtMiddleware() gin.HandlerFunc {
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

func loggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//开始时间
		startTime := time.Now()
		//处理请求
		c.Next()
		//结束时间
		endTime := time.Now()
		// 执行时间
		latencyTime := endTime.Sub(startTime).Milliseconds()
		//请求方式
		reqMethod := c.Request.Method
		//请求路由
		reqUrl := c.Request.RequestURI
		//状态码
		statusCode := c.Writer.Status()
		//请求ip
		clientIP := c.ClientIP()

		// 日志格式
		log.Log.Logger.WithFields(logrus.Fields{
			"status_code":  statusCode,
			"latency_time": latencyTime,
			"client_ip":    clientIP,
			"req_method":   reqMethod,
			"req_uri":      reqUrl,
		}).Info()
	}
}

// http服务包
package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	http_admin "registeruser/app/admin/server/http"
	"registeruser/conf/global"
	"registeruser/conf/log"
	"registeruser/util/ginMiddleware"
	"time"
)

func router(r *gin.Engine) *gin.Engine {
	r.Use(ginMiddleware.LoggerMiddleware())
	r.Use(ginMiddleware.CorsMiddleware())
	r.Use(ginMiddleware.RecoverMiddleware())
	http_admin.Register(r)
	return r
}

// 运行http服务
func Run() {
	gin.DefaultWriter = io.MultiWriter(log.Log.Logger.Writer(), os.Stdout)
	//gin.DefaultWriter = log.Log.Logger.Writer()
	router := router(gin.Default())
	address := fmt.Sprintf(":%s", global.CONFIG.App.Addr)
	server := &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(server.ListenAndServe())
}

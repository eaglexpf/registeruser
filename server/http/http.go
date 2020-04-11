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
	//r.Use(ginMiddleware.JWTMiddleware())
	//r.GET("/", func(c *gin.Context) {
	//	//jwt := NewJWT()
	//	token, err := util.NewJWT().CreateToken(&global.JwtClaims{})
	//	c.JSON(200, gin.H{
	//		"code": 0,
	//		"msg":  token,
	//		"err":  err,
	//	})
	//})
	//r.GET("/index", func(c *gin.Context) {
	//	//jwt := NewJWT()
	//	sign := c.DefaultQuery("sign", "asdsadasdad")
	//	token, err := util.NewJWT().ParseToken(sign)
	//	c.JSON(200, gin.H{
	//		"code": 0,
	//		"msg":  token,
	//		"err":  err,
	//		"sign": sign,
	//	})
	//})
	http_admin.Register(r)
	return r
}

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

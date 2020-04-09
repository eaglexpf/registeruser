package init

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"registeruser/entity/global"
	"registeruser/log"
	"registeruser/middleware"
	"registeruser/util"
	"time"
)

func router() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.LoggerMiddleware())
	r.Use(middleware.CorsMiddleware())
	r.Use(middleware.JWTMiddleware())
	r.GET("/", func(c *gin.Context) {
		//jwt := NewJWT()
		token, err := util.NewJWT().CreateToken(&global.JwtClaims{})
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  token,
			"err":  err,
		})
	})
	r.GET("/index", func(c *gin.Context) {
		//jwt := NewJWT()
		sign := c.DefaultQuery("sign", "asdsadasdad")
		token, err := util.NewJWT().ParseToken(sign)
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  token,
			"err":  err,
			"sign": sign,
		})
	})
	return r
}

func Run() {
	router := router()
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

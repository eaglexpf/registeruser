package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"registeruser/entity/global"
	"registeruser/log"
	"time"
)

func router() *gin.Engine {
	r := gin.Default()
	//gin.SetMode("debug")
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "success",
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
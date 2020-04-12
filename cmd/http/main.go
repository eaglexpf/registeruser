// http程序入口
package main

import (
	"registeruser/conf"
	"registeruser/server/http"
)

// main 程序入口
func main() {
	defer conf.Unload()
	http.Run()
}

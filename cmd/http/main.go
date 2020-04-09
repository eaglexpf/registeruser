package main

import (
	"registeruser/conf"
	"registeruser/server/http"
)

func main() {
	defer conf.Unload()
	http.Run()
}

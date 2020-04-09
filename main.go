package main

import (
	"registeruser/app"
)

func main() {
	defer app.Unload()
	app.Run()
}

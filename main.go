package main

import (
	app "registeruser/init"
)

func main() {
	defer app.Unload()
	app.Run()
}

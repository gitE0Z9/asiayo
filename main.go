package main

import (
	"asiayo/app"
)

func main() {
	server := app.SetupRoute()
	server.Run()
}

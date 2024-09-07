package main

import (
	"asiayo/application"
)

func main() {
	server := application.SetupRoute()
	application.RunServer(server, ":8080")
}

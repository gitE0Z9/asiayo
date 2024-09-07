package application

import (
	"asiayo/exchange_rate"

	"github.com/gin-gonic/gin"
)

func SetupRoute() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		exchange_rate.RegisterToGroup(v1)
	}

	return router
}

func RunServer(server *gin.Engine, port string) {
	server.Run(port)
}

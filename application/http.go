package application

import (
	docs "asiayo/docs"
	"asiayo/exchange_rate"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoute() *gin.Engine {
	router := gin.Default()

	docs.SwaggerInfo.Title = "AsiaYo exchange rate API"
	docs.SwaggerInfo.Description = ""
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api/v1"
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	v1 := router.Group("/api/v1")
	{
		exchange_rate.RegisterToGroup(v1)
	}

	return router
}

func RunServer(server *gin.Engine, port string) {
	server.Run(port)
}

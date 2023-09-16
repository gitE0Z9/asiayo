package app

import "github.com/gin-gonic/gin"

func SetupRoute() *gin.Engine {
	router := gin.Default()
	router.GET("/exchange-rate", exchangeRateConversion)

	return router
}

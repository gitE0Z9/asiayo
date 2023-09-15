package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	app.GET("/exchange-rate", exchangeRateConversion)
	app.Run()
}

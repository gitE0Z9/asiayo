package exchange_rate

import "github.com/gin-gonic/gin"

func RegisterToGroup(rootGroup *gin.RouterGroup) {
	handler := NewExchangeRateHandler()

	group := rootGroup.Group("/exchange-rate")
	{
		// more restful than homework
		group.GET("/:source/conversion/:target", handler.Conversion)
	}
}

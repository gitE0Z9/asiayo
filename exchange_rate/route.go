package exchange_rate

import "github.com/gin-gonic/gin"

func RegisterToGroup(rootGroup *gin.RouterGroup) {
	group := rootGroup.Group("/exchange-rate")
	{
		group.GET("", Conversion)
	}
}

package exchange_rate

import "github.com/gin-gonic/gin"

func NewErrorResponse() gin.H {
	return gin.H{
		"msg":    "error",
		"amount": "0",
	}
}

package app

import (
	"math"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type ExchangeRateConversionQuery struct {
	Source string `form:"source" binding:"required"`
	Target string `form:"target" binding:"required"`
	Amount string `form:"amount" binding:"required"`
}

func exchangeRateConversion(c *gin.Context) {
	var query ExchangeRateConversionQuery
	var convertedAmount float64

	if c.ShouldBind(&query) == nil {
		exchangeRate, ok := EXCHANGE_RATE_TABLE[query.Source][query.Target]

		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg":    "error",
				"amount": "0",
			})

			return
		}

		r, _ := regexp.Compile(`([0-9,.]+)`)
		match := r.FindString(query.Amount)
		match = strings.Replace(match, ",", "", -1)
		amount, err := strconv.ParseFloat(match, 32)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg":    "error",
				"amount": "0",
			})

			return
		}

		convertedAmount = math.Round((exchangeRate * amount * 100)) / 100

		printer := message.NewPrinter(language.English)
		c.JSON(http.StatusOK, gin.H{
			"msg":    "success",
			"amount": printer.Sprintf("$%.2f", convertedAmount),
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":    "error",
			"amount": "0",
		})
	}
}

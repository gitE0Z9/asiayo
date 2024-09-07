package exchange_rate

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

func Conversion(ctx *gin.Context) {
	var query ExchangeRateConversionQuery
	if ctx.BindQuery(&query) != nil {
		ctx.JSON(http.StatusBadRequest, NewErrorResponse())
		return
	}

	// find exchange rate from table
	exchangeRate, ok := RateTable[query.Source][query.Target]
	if !ok {
		ctx.JSON(http.StatusBadRequest, NewErrorResponse())
		return
	}

	// parse amount from string to float32
	r, _ := regexp.Compile(`([0-9,.]+)`)
	match := r.FindString(query.Amount)
	match = strings.Replace(match, ",", "", -1)
	amount, err := strconv.ParseFloat(match, 32)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, NewErrorResponse())
		return
	}

	// conversion here
	convertedAmount := math.Round((exchangeRate * amount * 100)) / 100

	// write response with required format
	printer := message.NewPrinter(language.English)
	ctx.JSON(http.StatusOK, gin.H{
		"msg":    "success",
		"amount": printer.Sprintf("$%.2f", convertedAmount),
	})
}

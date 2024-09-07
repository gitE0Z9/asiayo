package exchange_rate

import (
	"fmt"
	"math"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// ExchangeRateConversion godoc
// @Summary exchange rate conversion
// @Description exchange rate conversion
// @Tags exchange rate
// @Produce json
// @Param source query string true "source currency" Enums(JPY, TWD, USD)
// @Param target query string true "target currency" Enums(JPY, TWD, USD)
// @Param amount query string true "amount" example("$1,111.05")
// @Success 200 {object} exchange_rate.ConversionResponse "Success"
// @Failure 400 {object} exchange_rate.ConversionResponse "Bad parameter"
// @Router /exchange-rate [get]
func Conversion(ctx *gin.Context) {
	var query ConversionQuery
	if err := ctx.BindQuery(&query); err != nil {
		ctx.JSON(http.StatusBadRequest, NewErrorResponse(fmt.Sprintf("Bad parameter {%v}", err.Error())))
		return
	}

	// find exchange rate from table
	exchangeRate, ok := RateTable[query.Source][query.Target]
	if !ok {
		ctx.JSON(http.StatusBadRequest, NewErrorResponse("Given currency is not supported"))
		return
	}

	// parse amount from string to float32
	r, _ := regexp.Compile(`^\$\d{1,3}(,\d{3})*(\.\d{2})?$`)
	match := r.FindString(query.Amount)
	if match == "" {
		ctx.JSON(http.StatusBadRequest, NewErrorResponse("Given amount is in invalid format"))
		return
	}
	match = strings.Replace(match, ",", "", -1)
	match = strings.Replace(match, "$", "", -1)

	amount, err := strconv.ParseFloat(match, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, NewErrorResponse("Given amount {%f} is not a number"))
		return
	}

	// conversion here
	// 100 means two decimal places
	convertedAmount := math.Round((exchangeRate * amount * 100)) / 100

	// write response with required format
	printer := message.NewPrinter(language.English)
	ctx.JSON(http.StatusOK, gin.H{
		"msg":    "success",
		"amount": printer.Sprintf("$%.2f", convertedAmount),
	})
}

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

type ExchangeRateHandler struct {
}

func NewExchangeRateHandler() ExchangeRateHandler {
	return ExchangeRateHandler{}
}

// ExchangeRateConversion godoc
// @Summary exchange rate conversion
// @Description exchange rate conversion
// @Tags exchange rate
// @Produce json
// @Param source path string true "source currency" Enums(JPY, TWD, USD)
// @Param target path string true "target currency" Enums(JPY, TWD, USD)
// @Param amount query string true "amount" example("$1,111.05")
// @Success 200 {object} exchange_rate.ConversionResponse "Success"
// @Failure 400 {object} exchange_rate.ConversionResponse "Bad parameter"
// @Router /v1/exchange-rate/{source}/conversion/{target} [get]
func (h *ExchangeRateHandler) Conversion(ctx *gin.Context) {
	var param ConversionParam
	var query ConversionQuery

	if err := ctx.ShouldBindUri(&param); err != nil {
		ctx.JSON(http.StatusBadRequest, NewConversionErrorResponse(fmt.Sprintf("Bad parameter {%v}", err.Error())))
		return
	}
	source, ok := ExchangeRateFromValue(param.Source)
	if !ok {
		ctx.JSON(http.StatusBadRequest, NewConversionErrorResponse("Given currency is not supported"))
		return
	}
	target, ok := ExchangeRateFromValue(param.Target)
	if !ok {
		ctx.JSON(http.StatusBadRequest, NewConversionErrorResponse("Given currency is not supported"))
		return
	}

	if err := ctx.ShouldBindQuery(&query); err != nil {
		ctx.JSON(http.StatusBadRequest, NewConversionErrorResponse(fmt.Sprintf("Bad parameter {%v}", err.Error())))
		return
	}

	// find exchange rate from table
	exchangeRate, ok := RateTable[source][target]
	if !ok {
		ctx.JSON(http.StatusBadRequest, NewConversionErrorResponse("Given currency is not supported"))
		return
	}

	// parse amount from string to float32
	r, _ := regexp.Compile(`^\$\d{1,3}(,\d{3})*(\.\d{2})?$`)
	match := r.FindString(query.Amount)
	if match == "" {
		ctx.JSON(http.StatusBadRequest, NewConversionErrorResponse("Given amount is not valid format"))
		return
	}
	match = strings.Replace(match, ",", "", -1)
	match = strings.Replace(match, "$", "", -1)

	amount, err := strconv.ParseFloat(match, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, NewConversionErrorResponse("Given amount {%f} is not a number"))
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

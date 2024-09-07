package exchange_rate

type ExchangeRateConversionQuery struct {
	Source string `form:"source" binding:"required"`
	Target string `form:"target" binding:"required"`
	Amount string `form:"amount" binding:"required"`
}

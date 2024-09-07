package exchange_rate

type ConversionQuery struct {
	Source string `form:"source" binding:"required"`
	Target string `form:"target" binding:"required"`
	// @Description The accepted format for the amount is a string that starts with a dollar sign,
	// @Description followed by a positive number with up to two decimal places and thousand separators (commas).
	Amount string `form:"amount" binding:"required"`
}

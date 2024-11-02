package exchange_rate

type ConversionParam struct {
	Source string `uri:"source" binding:"required"`
	Target string `uri:"target" binding:"required"`
}

type ConversionQuery struct {
	// @Description The accepted format for the amount is a string that starts with a dollar sign,
	// @Description followed by a positive number with up to two decimal places and thousand separators (commas).
	Amount string `form:"amount" binding:"required"`
}

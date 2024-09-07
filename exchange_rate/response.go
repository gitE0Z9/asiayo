package exchange_rate

func NewErrorResponse(msg string) ConversionResponse {
	return ConversionResponse{
		Msg:    msg,
		Amount: "0",
	}
}

type ConversionResponse struct {
	Msg    string
	Amount string
}

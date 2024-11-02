package exchange_rate

func NewConversionErrorResponse(msg string) ConversionResponse {
	return ConversionResponse{
		Msg:    msg,
		Amount: "0",
	}
}

type ConversionResponse struct {
	Msg    string `json:"msg"`
	Amount string `json:"amount"`
}

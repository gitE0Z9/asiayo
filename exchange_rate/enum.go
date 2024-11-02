package exchange_rate

type ExchangeRate struct {
	Code  int
	Value string
}

var ExchangeRates = struct {
	JPY ExchangeRate
	USD ExchangeRate
	TWD ExchangeRate
}{
	JPY: ExchangeRate{
		Code:  1,
		Value: "JPY",
	},
	USD: ExchangeRate{
		Code:  2,
		Value: "USD",
	},
	TWD: ExchangeRate{
		Code:  3,
		Value: "TWD",
	},
}

func ExchangeRateFromValue(v string) (ExchangeRate, bool) {
	switch v {
	case "JPY":
		return ExchangeRates.JPY, true
	case "TWD":
		return ExchangeRates.TWD, true
	case "USD":
		return ExchangeRates.USD, true
	}

	return ExchangeRate{}, false
}

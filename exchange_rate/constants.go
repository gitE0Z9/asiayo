package exchange_rate

var RateTable = map[ExchangeRate]map[ExchangeRate]float64{
	ExchangeRates.TWD: {
		ExchangeRates.TWD: 1,
		ExchangeRates.JPY: 3.669,
		ExchangeRates.USD: 0.03281,
	},
	ExchangeRates.JPY: {
		ExchangeRates.TWD: 0.26956,
		ExchangeRates.JPY: 1,
		ExchangeRates.USD: 0.00885,
	},
	ExchangeRates.USD: {
		ExchangeRates.TWD: 30.444,
		ExchangeRates.JPY: 111.801,
		ExchangeRates.USD: 1,
	},
}

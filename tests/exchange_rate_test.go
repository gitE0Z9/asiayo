package tests

import (
	"asiayo/application"
	"asiayo/exchange_rate"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

const conversionPath = "/api/v1/exchange-rate/%s/conversion/%s"

func TestConversion(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		// setup
		router := application.SetupRoute()
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(
			http.MethodGet,
			fmt.Sprintf(conversionPath, "USD", "JPY"),
			nil,
		)
		q := req.URL.Query()
		q.Add("amount", "$1,525")
		req.URL.RawQuery = q.Encode()

		// action
		router.ServeHTTP(w, req)

		// assert
		var response exchange_rate.ConversionResponse
		json.Unmarshal(w.Body.Bytes(), &response)
		expectedResponse := exchange_rate.ConversionResponse{
			Msg:    "success",
			Amount: "$170,496.53",
		}
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, expectedResponse, response)
	})

	t.Run("EmptyField", func(t *testing.T) {
		tests := []struct {
			source   string
			target   string
			amount   string
			wantCode int
		}{
			{
				source:   "",
				target:   exchange_rate.ExchangeRates.JPY.Value,
				amount:   "$1,525",
				wantCode: http.StatusBadRequest,
			},
			{
				source: exchange_rate.ExchangeRates.USD.Value,
				target: "",
				amount: "$1,525",
				// XXX: because /USD/conversion/?amount=$123 is not matched
				wantCode: http.StatusNotFound,
			},
			{
				source:   exchange_rate.ExchangeRates.USD.Value,
				target:   exchange_rate.ExchangeRates.JPY.Value,
				amount:   "",
				wantCode: http.StatusBadRequest,
			},
		}

		for _, test := range tests {
			// setup
			router := application.SetupRoute()
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(
				http.MethodGet,
				fmt.Sprintf(conversionPath, test.source, test.target),
				nil,
			)
			q := req.URL.Query()
			q.Add("amount", test.amount)
			req.URL.RawQuery = q.Encode()

			// action
			router.ServeHTTP(w, req)

			// assert
			var response exchange_rate.ConversionResponse
			json.Unmarshal(w.Body.Bytes(), &response)
			// expectedResponse := exchange_rate.NewConversionErrorResponse("error")
			assert.Equal(t, test.wantCode, w.Code)
			// assert.Equal(t, expectedResponse, response)
		}
	})

	t.Run("OutOfCurrency", func(t *testing.T) {
		tests := []struct {
			source string
			target string
			amount string
		}{
			{
				source: "AAA",
				target: exchange_rate.ExchangeRates.JPY.Value,
				amount: "$1,525",
			},
			{
				source: exchange_rate.ExchangeRates.USD.Value,
				target: "AAA",
				amount: "$1,525",
			},
		}

		for _, test := range tests {
			// setup
			router := application.SetupRoute()
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(
				http.MethodGet,
				fmt.Sprintf(conversionPath, test.source, test.target),
				nil,
			)
			q := req.URL.Query()
			q.Add("amount", test.amount)
			req.URL.RawQuery = q.Encode()

			// action
			router.ServeHTTP(w, req)

			// assert
			var response exchange_rate.ConversionResponse
			json.Unmarshal(w.Body.Bytes(), &response)
			// expectedResponse := exchange_rate.NewConversionErrorResponse("error")
			assert.Equal(t, http.StatusBadRequest, w.Code)
			// assert.Equal(t, expectedResponse, response)
		}
	})

	t.Run("IncorrectFormatAmount", func(t *testing.T) {
		// setup
		router := application.SetupRoute()
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(
			http.MethodGet,
			fmt.Sprintf(conversionPath, exchange_rate.ExchangeRates.USD.Value, exchange_rate.ExchangeRates.JPY.Value),
			nil,
		)
		q := req.URL.Query()
		q.Add("amount", "1.2.3")
		req.URL.RawQuery = q.Encode()

		// action
		router.ServeHTTP(w, req)

		// assert
		var response exchange_rate.ConversionResponse
		json.Unmarshal(w.Body.Bytes(), &response)
		// expectedResponse := exchange_rate.NewConversionErrorResponse("error")
		assert.Equal(t, http.StatusBadRequest, w.Code)
		// assert.Equal(t, expectedResponse, response)
	})
}

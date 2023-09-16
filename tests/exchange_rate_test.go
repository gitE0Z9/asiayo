package tests

import (
	"asiayo/app"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuccess(t *testing.T) {
	router := app.SetupRoute()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/exchange-rate", nil)

	q := req.URL.Query()
	q.Add("source", "USD")
	q.Add("target", "JPY")
	q.Add("amount", "$1,525")
	req.URL.RawQuery = q.Encode()
	router.ServeHTTP(w, req)

	expectedResponse := map[string]string{
		"msg":    "success",
		"amount": "$170,496.53",
	}
	expectedResponseString, _ := json.Marshal(expectedResponse)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, string(expectedResponseString), w.Body.String())
}

func TestEmptySource(t *testing.T) {
	router := app.SetupRoute()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/exchange-rate", nil)

	q := req.URL.Query()
	q.Add("source", "")
	q.Add("target", "JPY")
	q.Add("amount", "$1,525")
	req.URL.RawQuery = q.Encode()
	router.ServeHTTP(w, req)

	expectedResponse := map[string]string{
		"msg":    "error",
		"amount": "0",
	}
	expectedResponseString, _ := json.Marshal(expectedResponse)

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, string(expectedResponseString), w.Body.String())
}

func TestEmptyTarget(t *testing.T) {
	router := app.SetupRoute()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/exchange-rate", nil)

	q := req.URL.Query()
	q.Add("source", "USD")
	q.Add("target", "")
	q.Add("amount", "$1,525")
	req.URL.RawQuery = q.Encode()
	router.ServeHTTP(w, req)

	expectedResponse := map[string]string{
		"msg":    "error",
		"amount": "0",
	}
	expectedResponseString, _ := json.Marshal(expectedResponse)

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, string(expectedResponseString), w.Body.String())
}

func TestEmptyAmount(t *testing.T) {
	router := app.SetupRoute()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/exchange-rate", nil)

	q := req.URL.Query()
	q.Add("source", "USD")
	q.Add("target", "JPY")
	q.Add("amount", "")
	req.URL.RawQuery = q.Encode()
	router.ServeHTTP(w, req)

	expectedResponse := map[string]string{
		"msg":    "error",
		"amount": "0",
	}
	expectedResponseString, _ := json.Marshal(expectedResponse)

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, string(expectedResponseString), w.Body.String())
}

func TestOutOfKeySource(t *testing.T) {
	router := app.SetupRoute()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/exchange-rate", nil)

	q := req.URL.Query()
	q.Add("source", "AAA")
	q.Add("target", "JPY")
	q.Add("amount", "$1,525")
	req.URL.RawQuery = q.Encode()
	router.ServeHTTP(w, req)

	expectedResponse := map[string]string{
		"msg":    "error",
		"amount": "0",
	}
	expectedResponseString, _ := json.Marshal(expectedResponse)

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, string(expectedResponseString), w.Body.String())
}

func TestOutOfKeyTarget(t *testing.T) {
	router := app.SetupRoute()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/exchange-rate", nil)

	q := req.URL.Query()
	q.Add("source", "USD")
	q.Add("target", "AAA")
	q.Add("amount", "$1,525")
	req.URL.RawQuery = q.Encode()
	router.ServeHTTP(w, req)

	expectedResponse := map[string]string{
		"msg":    "error",
		"amount": "0",
	}
	expectedResponseString, _ := json.Marshal(expectedResponse)

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, string(expectedResponseString), w.Body.String())
}

func TestIncorrectFormatAmount(t *testing.T) {
	router := app.SetupRoute()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/exchange-rate", nil)

	q := req.URL.Query()
	q.Add("source", "USD")
	q.Add("target", "JPY")
	q.Add("amount", "1.2.3")
	req.URL.RawQuery = q.Encode()
	router.ServeHTTP(w, req)

	expectedResponse := map[string]string{
		"msg":    "error",
		"amount": "0",
	}
	expectedResponseString, _ := json.Marshal(expectedResponse)

	assert.Equal(t, 500, w.Code)
	assert.Equal(t, string(expectedResponseString), w.Body.String())
}

package tests

import (
	"asiayo/application"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

const conversionPath = "/api/v1/exchange-rate"

func TestSuccess(t *testing.T) {
	router := application.SetupRoute()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", conversionPath, nil)

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

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, string(expectedResponseString), w.Body.String())
}

func TestEmptySource(t *testing.T) {
	router := application.SetupRoute()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", conversionPath, nil)

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

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, string(expectedResponseString), w.Body.String())
}

func TestEmptyTarget(t *testing.T) {
	router := application.SetupRoute()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", conversionPath, nil)

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

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, string(expectedResponseString), w.Body.String())
}

func TestEmptyAmount(t *testing.T) {
	router := application.SetupRoute()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", conversionPath, nil)

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

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, string(expectedResponseString), w.Body.String())
}

func TestOutOfKeySource(t *testing.T) {
	router := application.SetupRoute()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", conversionPath, nil)

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

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, string(expectedResponseString), w.Body.String())
}

func TestOutOfKeyTarget(t *testing.T) {
	router := application.SetupRoute()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", conversionPath, nil)

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

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, string(expectedResponseString), w.Body.String())
}

func TestIncorrectFormatAmount(t *testing.T) {
	router := application.SetupRoute()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", conversionPath, nil)

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

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, string(expectedResponseString), w.Body.String())
}

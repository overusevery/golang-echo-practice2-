package testutil

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/labstack/echo/v4"
)

func Post(e *echo.Echo, url string, jsonPath string) *httptest.ResponseRecorder {
	requestJson, err := os.ReadFile(jsonPath)
	if err != nil {
		panic(err)
	}
	req := httptest.NewRequest(http.MethodPost, url, bytes.NewReader(requestJson))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	e.ServeHTTP(res, req)
	return res
}

func GET(e *echo.Echo, url string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(http.MethodGet, url, nil)
	res := httptest.NewRecorder()
	e.ServeHTTP(res, req)
	return res
}

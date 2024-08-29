package testutil

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/labstack/echo/v4"
)

// Data
//
//	{
//		"sub": "11111111-1111-1111-1111-111111111111",
//		"iss": "someiss",
//		"client_id": "someclient_id",
//		"scope": "mybackendapi/getcustomer mybackendapi/editcustomer",
//		"exp": 1824767332,
//		"iat": 1724763732,
//		"jti": "22222222-2222-2222-2222-222222222222"
//	}
var authToken = "Bearer eyJraWQiOiJzb21la2lkIiwiYWxnIjoiSFMyNTYifQ.eyJzdWIiOiIxMTExMTExMS0xMTExLTExMTEtMTExMS0xMTExMTExMTExMTEiLCJpc3MiOiJzb21laXNzIiwiY2xpZW50X2lkIjoic29tZWNsaWVudF9pZCIsInNjb3BlIjoibXliYWNrZW5kYXBpL2dldGN1c3RvbWVyIG15YmFja2VuZGFwaS9lZGl0Y3VzdG9tZXIiLCJleHAiOjE4MjQ3NjczMzIsImlhdCI6MTcyNDc2MzczMiwianRpIjoiMjIyMjIyMjItMjIyMi0yMjIyLTIyMjItMjIyMjIyMjIyMjIyIn0.AUPdh5v9fvna4U8NiRKK5aq4AgFzwu1WAMwKC7FSiCY" //nolint:gosec,lll, this is just example dummy token

// Data
//
//	{
//		"sub": "11111111-1111-1111-1111-111111111111",
//		"iss": "someiss",
//		"client_id": "someclient_id",
//		"scope": "mybackendapi/notexisting",
//		"exp": 1824767332,
//		"iat": 1724763732,
//		"jti": "22222222-2222-2222-2222-222222222222"
//	}
var AuthTokenScopeUnmatch = "Bearer eyJraWQiOiJzb21la2lkIiwiYWxnIjoiSFMyNTYifQ.eyJzdWIiOiIxMTExMTExMS0xMTExLTExMTEtMTExMS0xMTExMTExMTExMTEiLCJpc3MiOiJzb21laXNzIiwiY2xpZW50X2lkIjoic29tZWNsaWVudF9pZCIsInNjb3BlIjoibXliYWNrZW5kYXBpL25vdGV4aXN0aW5nIiwiZXhwIjoxODI0NzY3MzMyLCJpYXQiOjE3MjQ3NjM3MzIsImp0aSI6IjIyMjIyMjIyLTIyMjItMjIyMi0yMjIyLTIyMjIyMjIyMjIyMiJ9.e1-BOmvpCbTcuBIlQ3qCl58fsaRQL6bNXUjW_MdXR2M" //nolint:gosec,lll, this is just example dummy token

type option func(*http.Request)

func WithAuthToken(token string) option {
	return func(req *http.Request) {
		req.Header.Set("Authorization", token)
	}
}

func Post(e *echo.Echo, url string, jsonPath string) *httptest.ResponseRecorder {
	requestJson, err := os.ReadFile(jsonPath)
	if err != nil {
		panic(err)
	}
	req := httptest.NewRequest(http.MethodPost, url, bytes.NewReader(requestJson))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", authToken)
	res := httptest.NewRecorder()
	e.ServeHTTP(res, req)
	return res
}

func GET(e *echo.Echo, url string, opts ...option) *httptest.ResponseRecorder {
	req := httptest.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("Authorization", authToken)

	for _, opt := range opts {
		opt(req)
	}

	res := httptest.NewRecorder()
	e.ServeHTTP(res, req)
	return res
}

func PUT(e *echo.Echo, url string, jsonPath string) *httptest.ResponseRecorder {
	requestJson, err := os.ReadFile(jsonPath)
	if err != nil {
		panic(err)
	}
	req := httptest.NewRequest(http.MethodPut, url, bytes.NewReader(requestJson))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", authToken)
	res := httptest.NewRecorder()
	e.ServeHTTP(res, req)
	return res
}

func DELETE(e *echo.Echo, url string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(http.MethodDelete, url, nil)
	req.Header.Set("Authorization", authToken)
	res := httptest.NewRecorder()
	e.ServeHTTP(res, req)
	return res
}

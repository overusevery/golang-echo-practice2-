package customemiddleware

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// HEADER
//
//	{
//		"kid": "somekid",
//		"alg": "HS256"
//	}
//
// DATA
//
//	{
//		"sub": "11111111-1111-1111-1111-111111111111",
//		"iss": "someiss",
//		"client_id": "someclient_id",
//		"scope": "mybackendapi/my_scope_name",
//		"exp": 1824767332,
//		"iat": 1724763732,
//		"jti": "22222222-2222-2222-2222-222222222222"
//	  }
var example_token = "eyJraWQiOiJzb21la2lkIiwiYWxnIjoiSFMyNTYifQ.eyJzdWIiOiIxMTExMTExMS0xMTExLTExMTEtMTExMS0xMTExMTExMTExMTEiLCJpc3MiOiJzb21laXNzIiwiY2xpZW50X2lkIjoic29tZWNsaWVudF9pZCIsInNjb3BlIjoibXliYWNrZW5kYXBpL215X3Njb3BlX25hbWUiLCJleHAiOjE4MjQ3NjczMzIsImlhdCI6MTcyNDc2MzczMiwianRpIjoiMjIyMjIyMjItMjIyMi0yMjIyLTIyMjItMjIyMjIyMjIyMjIyIn0.Pc6JBn6HHDd2MWJkne1SroqWaxFkIySW65KM1X83XNE"

// DATA
//
//	{
//		"sub": "11111111-1111-1111-1111-111111111111",
//		"iss": "someiss",
//		"client_id": "someclient_id",
//		"scope": "mybackendapi/A mybackendapi/B",
//		"exp": 1824767332,
//		"iat": 1724763732,
//		"jti": "22222222-2222-2222-2222-222222222222"
//	  }
var example_token_multi_scope = "eyJraWQiOiJzb21la2lkIiwiYWxnIjoiSFMyNTYifQ.eyJzdWIiOiIxMTExMTExMS0xMTExLTExMTEtMTExMS0xMTExMTExMTExMTEiLCJpc3MiOiJzb21laXNzIiwiY2xpZW50X2lkIjoic29tZWNsaWVudF9pZCIsInNjb3BlIjoibXliYWNrZW5kYXBpL0EgbXliYWNrZW5kYXBpL0IiLCJleHAiOjE4MjQ3NjczMzIsImlhdCI6MTcyNDc2MzczMiwianRpIjoiMjIyMjIyMjItMjIyMi0yMjIyLTIyMjItMjIyMjIyMjIyMjIyIn0.HssGOJjhhfuc9oUQkLz5PvyD5QYBk63h7nfKTTu3mw4"

func TestParseAuthorizationToken(t *testing.T) {
	t.Run("single scope", func(t *testing.T) {
		e := echo.New()
		hw := []byte("Hello, World!")
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(hw))
		res := httptest.NewRecorder()
		c := e.NewContext(req, res)
		var recievedContext echo.Context
		dummyHandler := func(c echo.Context) error {
			recievedContext = c
			body, err := io.ReadAll(c.Request().Body)
			if err != nil {
				return err
			}
			return c.String(http.StatusOK, string(body))
		}

		req.Header.Set(echo.HeaderAuthorization, "Bearer "+example_token)
		ParseAuthorizationToken(dummyHandler)(c)

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, hw, res.Body.Bytes())
		assert.Equal(t, "11111111-1111-1111-1111-111111111111", extract(recievedContext, "user_id"))
		assert.Equal(t, []string{"mybackendapi/my_scope_name"}, extract(recievedContext, "scope"))
	})

	t.Run("multi scope", func(t *testing.T) {
		e := echo.New()
		hw := []byte("Hello, World!")
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(hw))
		res := httptest.NewRecorder()
		c := e.NewContext(req, res)
		var recievedContext echo.Context
		dummyHandler := func(c echo.Context) error {
			recievedContext = c
			body, err := io.ReadAll(c.Request().Body)
			if err != nil {
				return err
			}
			return c.String(http.StatusOK, string(body))
		}

		req.Header.Set(echo.HeaderAuthorization, "Bearer "+example_token_multi_scope)
		ParseAuthorizationToken(dummyHandler)(c)

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, hw, res.Body.Bytes())
		assert.Equal(t, "11111111-1111-1111-1111-111111111111", extract(recievedContext, USER_ID))
		assert.Equal(t, []string{"mybackendapi/A", "mybackendapi/B"}, extract(recievedContext, SCOPE))
	})
	t.Run("request without access token is invalid", func(t *testing.T) {
		e := echo.New()
		hw := []byte("Hello, World!")
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(hw))
		res := httptest.NewRecorder()
		c := e.NewContext(req, res)
		var recievedContext echo.Context
		dummyHandler := func(c echo.Context) error {
			require.Fail(t, "should not be called")
			recievedContext = c
			body, err := io.ReadAll(c.Request().Body)
			if err != nil {
				return err
			}
			return c.String(http.StatusOK, string(body))
		}

		// no token
		// req.Header.Set(echo.HeaderAuthorization, "Bearer "+example_token)
		ParseAuthorizationToken(dummyHandler)(c)

		assert.Equal(t, http.StatusUnauthorized, res.Code)
		assert.NotEqual(t, hw, res.Body.Bytes())
		assert.Equal(t, nil, recievedContext)
	})
	t.Run("non access token value is invalid", func(t *testing.T) {
		e := echo.New()
		hw := []byte("Hello, World!")
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(hw))
		res := httptest.NewRecorder()
		c := e.NewContext(req, res)
		var recievedContext echo.Context
		dummyHandler := func(c echo.Context) error {
			require.Fail(t, "should not be called")
			recievedContext = c
			body, err := io.ReadAll(c.Request().Body)
			if err != nil {
				return err
			}
			return c.String(http.StatusOK, string(body))
		}

		req.Header.Set(echo.HeaderAuthorization, "Bearer "+"xxx")
		ParseAuthorizationToken(dummyHandler)(c)

		assert.Equal(t, http.StatusUnauthorized, res.Code)
		assert.NotEqual(t, hw, res.Body.Bytes())
		assert.Equal(t, nil, recievedContext)
	})
	t.Run("token signature is invalid", func(t *testing.T) {
		e := echo.New()
		hw := []byte("Hello, World!")
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(hw))
		res := httptest.NewRecorder()
		c := e.NewContext(req, res)
		var recievedContext echo.Context
		dummyHandler := func(c echo.Context) error {
			require.Fail(t, "should not be called")
			recievedContext = c
			body, err := io.ReadAll(c.Request().Body)
			if err != nil {
				return err
			}
			return c.String(http.StatusOK, string(body))
		}

		req.Header.Set(echo.HeaderAuthorization, "Bearer "+"eyJraWQiOiJzb21la2lkIiwiYWxnIjoiSFMyNTYifQ.eyJzdWIiOiIxMTExMTExMS0xMTExLTExMTEtMTExMS0xMTExMTExMTExMTEiLCJpc3MiOiJzb21laXNzIiwiY2xpZW50X2lkIjoic29tZWNsaWVudF9pZCIsInNjb3BlIjoibXliYWNrZW5kYXBpL0EgbXliYWNrZW5kYXBpL0IiLCJleHAiOjE4MjQ3NjczMzIsImlhdCI6MTcyNDc2MzczMiwianRpIjoiMjIyMjIyMjItMjIyMi0yMjIyLTIyMjItMjIyMjIyMjIyMjIyIn0.XXXXXXXXXX")
		ParseAuthorizationToken(dummyHandler)(c)

		assert.Equal(t, http.StatusUnauthorized, res.Code)
		assert.NotEqual(t, hw, res.Body.Bytes())
		assert.Equal(t, nil, recievedContext)
	})
	t.Run("token signature should not be empty", func(t *testing.T) {
		e := echo.New()
		hw := []byte("Hello, World!")
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(hw))
		res := httptest.NewRecorder()
		c := e.NewContext(req, res)
		var recievedContext echo.Context
		dummyHandler := func(c echo.Context) error {
			require.Fail(t, "should not be called")
			recievedContext = c
			body, err := io.ReadAll(c.Request().Body)
			if err != nil {
				return err
			}
			return c.String(http.StatusOK, string(body))
		}

		req.Header.Set(echo.HeaderAuthorization, "Bearer "+"eyJraWQiOiJzb21la2lkIiwiYWxnIjoiSFMyNTYifQ.eyJzdWIiOiIxMTExMTExMS0xMTExLTExMTEtMTExMS0xMTExMTExMTExMTEiLCJpc3MiOiJzb21laXNzIiwiY2xpZW50X2lkIjoic29tZWNsaWVudF9pZCIsInNjb3BlIjoibXliYWNrZW5kYXBpL0EgbXliYWNrZW5kYXBpL0IiLCJleHAiOjE4MjQ3NjczMzIsImlhdCI6MTcyNDc2MzczMiwianRpIjoiMjIyMjIyMjItMjIyMi0yMjIyLTIyMjItMjIyMjIyMjIyMjIyIn0")
		ParseAuthorizationToken(dummyHandler)(c)

		assert.Equal(t, http.StatusUnauthorized, res.Code)
		assert.NotEqual(t, hw, res.Body.Bytes())
		assert.Equal(t, nil, recievedContext)
	})
}

func extract(ctx echo.Context, key string) interface{} {
	return ctx.Request().Context().Value(key)
}

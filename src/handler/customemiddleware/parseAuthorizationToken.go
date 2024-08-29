package customemiddleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	openapi "github.com/overusevery/golang-echo-practice2/src/handler/openapigenmodel/go"
)

var (
	USER_ID = "user_id"
	SCOPE   = "scope"
)

func ParseAuthorizationToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		//In production code, using kong or api gateway to verify jwt may be easier.
		claims, isInvalid := parseAuthorizationToken(c)
		if isInvalid {
			return c.JSON(http.StatusUnauthorized, openapi.Error{
				Message: "valid API key is not provided",
			})
		}

		setValueToContext(c, USER_ID, claims.Subject)
		setValueToContext(c, SCOPE, strings.Fields(claims.Scope))
		return next(c)
	}
}
func setValueToContext(c echo.Context, key string, value any) {
	c.SetRequest(c.Request().WithContext(context.WithValue(c.Request().Context(), key, value)))
}

func parseAuthorizationToken(c echo.Context) (*CustomClaims, bool) {
	authToken := c.Request().Header.Get("Authorization")
	if !strings.HasPrefix(authToken, "Bearer ") {
		return nil, true
	}

	bearerToken := strings.TrimPrefix(authToken, "Bearer ")
	claims := &CustomClaims{}

	_, err := jwt.ParseWithClaims(bearerToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("mysecrets"), nil
	})
	if err != nil {
		return nil, true
	}
	//if access token's verification is needed in this server,  check should be exected at here like this
	// if claims.Issuer != xxx{
	// 	return some err
	// }
	// if claims.Audience != xx {
	// 	return some err
	// }
	// etc...
	return claims, false
}

type CustomClaims struct {
	Scope string `json:"scope"`
	jwt.RegisteredClaims
}

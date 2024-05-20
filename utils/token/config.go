package token

import (
	"errors"
	"net/http"
	"os"

	"github.com/google/uuid"
	msg "github.com/satryanararya/go-chefbot/constants/message"
	http_util "github.com/satryanararya/go-chefbot/utils/http"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type JWTClaim struct {
	ID uuid.UUID `json:"id"`
	jwt.RegisteredClaims
}

func GetJWTConfig() echojwt.Config {
	return echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return &JWTClaim{}
		},
		ErrorHandler: jwtErrorHandler,
		SigningKey:   []byte(os.Getenv("JWT_SECRET")),
	}
}

func jwtErrorHandler(c echo.Context, err error) error {
	if errors.Is(err, echojwt.ErrJWTInvalid) {
		return http_util.HandleErrorResponse(c, http.StatusUnauthorized, msg.MsgInvalidToken)
	}

	return http_util.HandleErrorResponse(c, http.StatusInternalServerError, msg.MsgUnauthorized)
}
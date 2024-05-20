package token

import (
	"os"
	"time"

	"github.com/google/uuid"
	error_util "github.com/satryanararya/go-chefbot/utils/error"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type TokenUtil interface {
	GenerateToken(id uuid.UUID) (string, error)
	GetClaims(c echo.Context) *JWTClaim
}

type tokenUtil struct {}

func NewTokenUtil() *tokenUtil {
	return &tokenUtil{}
}

func (tu *tokenUtil) GenerateToken(id uuid.UUID) (string, error) {
	claims := JWTClaim{
		ID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 12)),
		},
	}
	unsignedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := unsignedToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", error_util.ErrFailedGeneratingToken
	}
	return signedToken, nil
}

func (tu *tokenUtil) GetClaims(c echo.Context) *JWTClaim {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*JWTClaim)
	return claims
}
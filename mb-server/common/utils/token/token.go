package token

import (
	"mb-server/common/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func CreateJwt(email string) (string, *CustomClaims, error) {
	expireAt := time.Now().Add(24 * 3 * time.Hour).Unix()
	claim := &CustomClaims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Unix(expireAt, 0)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	secretKey := config.GetWebConfig().SecretKey
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenStr, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", nil, err
	}
	return tokenStr, claim, nil

}

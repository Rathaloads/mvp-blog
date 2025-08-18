package token

import (
	"fmt"
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

func ValidateToken(tokenStr string) (*CustomClaims, error) {
	var claim CustomClaims
	token, err := jwt.ParseWithClaims(tokenStr, &claim, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		secretKey := config.GetWebConfig().SecretKey
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid {
		return &claim, nil
	}
	return nil, fmt.Errorf("token unvaild")

}

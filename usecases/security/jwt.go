package security

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func Encode(appSecret string, payload jwt.MapClaims, expireAt *time.Time) (string, error) {
	if expireAt != nil {
		payload["exp"] = expireAt.Unix()
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return token.SignedString([]byte(appSecret))
}

func Decode(appSecret string, tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(appSecret), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	var ve *jwt.ValidationError
	if errors.As(err, &ve) {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return nil, fmt.Errorf("that's not even a token")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return nil, fmt.Errorf("timing is everything")
		}
	}
	return nil, fmt.Errorf("couldn't handle this token: %w", err)
}

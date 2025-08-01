package util

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var (
	key []byte
)

type Claims struct {
	Uid int64 `json:"uid"`
	jwt.StandardClaims
}

func init() {
	key = []byte("q3Ga3OVFED7wWySV")
}

func JwtEncode(c Claims, expire int64, keys []byte) (string, error) {
	if c.ExpiresAt == 0 {
		c.ExpiresAt = time.Now().Unix() + expire
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// Sign and get the complete encoded token as a string using the secret
	if keys != nil {
		key = keys
	}
	return token.SignedString(key)
}

func JwtDecode(s string, keys []byte) (*Claims, error) {
	var err error
	// sample token is expired.  override time so it parses as valid
	if keys != nil {
		key = keys
	}
	if s == "" {
		return &Claims{}, errors.New("token不能为空")
	}
	token, err := jwt.ParseWithClaims(s, &Claims{}, func(token *jwt.Token) (i interface{}, err error) {
		return key, nil
	})
	if err != nil {
		return &Claims{}, err
	}

	if !token.Valid {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				err = errors.New("that's not even a token")
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				// Token is either expired or not active yet
				err = errors.New("timing is everything")
			} else {
				err = errors.New("couldn't handle this token")
			}
		} else {
			err = errors.New("couldn't handle this token")
		}
		return &Claims{}, err
	}
	claims, ok := token.Claims.(*Claims)
	if !ok {
		return &Claims{}, errors.New("Couldn't handle this token:")
	}

	return claims, nil
}

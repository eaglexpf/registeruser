package util

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"registeruser/conf/global"
	"time"
)

var (
	TokenExpired     error = errors.New("Token is expired")
	TokenNotValidYet error = errors.New("Token not active yet")
	TokenMalformed   error = errors.New("That's not even a token")
	TokenInvalid     error = errors.New("Couldn't handle this token:")
)

func NewJWT() *appJWT {
	return &appJWT{
		signKey: []byte(global.CONFIG.Jwt.Sign),
	}
}

type appJWT struct {
	signKey []byte
}

func (j *appJWT) CreateToken(claims *global.JwtClaims) (string, error) {
	claims.StandardClaims.ExpiresAt = time.Now().Add(time.Duration(global.CONFIG.Jwt.Express) * time.Second).Unix()
	claims.StandardClaims.NotBefore = time.Now().Unix()
	claims.StandardClaims.Issuer = global.CONFIG.Jwt.Issuer
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.signKey)
}

func (j *appJWT) ParseToken(token string) (*global.JwtClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &global.JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.signKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*global.JwtClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, TokenInvalid
}

func (j *appJWT) RefreshToken(token string) (string, error) {
	tokenCliams, err := j.ParseToken(token)
	if err != nil {
		return "", err
	}
	return j.CreateToken(tokenCliams)
}

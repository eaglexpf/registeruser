package global

import "github.com/dgrijalva/jwt-go"

type JwtClaims struct {
	jwt.StandardClaims
}

package global

import "github.com/dgrijalva/jwt-go"

// jwt数据格式
type JwtClaims struct {
	jwt.StandardClaims
	UUID string
}

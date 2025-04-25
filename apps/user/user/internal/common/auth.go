package common

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

func GetJwtToken(secret string, seconds int64, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	iat := time.Now().Unix()
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["jwtUserId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secret))
}

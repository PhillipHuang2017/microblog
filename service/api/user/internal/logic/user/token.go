package logic

import "github.com/dgrijalva/jwt-go"

func getToken(secret string, iat, expire int64, uid string) (string, error)  {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + expire
	claims["iat"] = iat
	claims["uid"] = uid
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
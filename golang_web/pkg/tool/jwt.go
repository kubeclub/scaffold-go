package tool

import (
	"time"
	models "kubeclub/scaffold_go_web/dao"
	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("a_secret_creat")

type CustomClaims struct {
	ID         int
	Username   string
	Role       string
	BufferTime int64
	jwt.RegisteredClaims
}

func ReleaseToken(u models.User) (string, error) {
	/*
		创建token
	*/
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &CustomClaims{
		ID:       u.Id,
		Username: u.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "luban",
			Subject:   "user token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseToken(token string) (*jwt.Token, *CustomClaims, error) {
	/*
		解析token
	*/
	claims := &CustomClaims{}
	tk, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return tk, claims, err
}

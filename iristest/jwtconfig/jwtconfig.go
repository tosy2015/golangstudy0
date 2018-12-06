package jwtconfig

import (
	"github.com/dgrijalva/jwt-go"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
)

const (
	Key = "RealSecret"
)
var (
	j = jwtmiddleware.New(jwtmiddleware.Config{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(Key), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
		//Expiration : true,
	})
)

func Get() *jwtmiddleware.Middleware{
	return j
}


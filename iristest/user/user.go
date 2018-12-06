package user

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/golangstudy0/iristest/jwtconfig"
	"github.com/kataras/iris"
	"log"
	"strconv"
	"time"
)

func New()*Handler {
	api := Handler{
		Login: Login,
		Todo: Todo,
	}
	return &api
}

type Handler struct {
	Login func(ctx iris.Context)
	Todo func(ctx iris.Context)
}

func Login(ctx iris.Context){
	log.Println("Login:call")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid": strconv.FormatInt(time.Now().UnixNano(),10),
		"exp": time.Now().Unix()+24*60*60,
	})
	tokenString, _ := token.SignedString([]byte(jwtconfig.Key))
	ctx.JSON(iris.Map{"Authorization": "Bearer "+tokenString})
	log.Println("Login:return ",tokenString)
}

func Todo(ctx iris.Context){
	log.Println("Todo:call")
	//TODO
}

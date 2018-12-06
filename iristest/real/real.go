package real

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/golangstudy0/iristest/jwtconfig"
	"github.com/kataras/iris"
	"log"
)

func New()*Handler {
	api := Handler{
		GetList: GetList,
		Todo: Todo,
	}
	return &api
}

type Handler struct {
	GetList func(ctx iris.Context)
	Todo func(ctx iris.Context)
}

func GetList(ctx iris.Context){
	log.Println("GetList:call")
	userToken := jwtconfig.Get().Get(ctx)

	if claims, ok := userToken.Claims.(jwt.MapClaims); ok && userToken.Valid {
		ctx.Writef("uid:%s\n",claims["uid"].(string))
	}
	user := ctx.Values().Get("jwt").(*jwt.Token)
	ctx.Writef("This is an authenticated request\n")
	ctx.Writef("Claim content:\n")
	ctx.Writef("%s", user.Signature)
	//TODO
}

func Todo(ctx iris.Context){
	log.Println("Todo:call")
	//TODO
}

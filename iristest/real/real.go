package real

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/golangstudy0/iristest/jwtconfig"
	"github.com/kataras/iris"
	"log"
	"reflect"
	"time"
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
	//userToken := ctx.Values().Get("jwt").(*jwt.Token)

	if claims, ok := userToken.Claims.(jwt.MapClaims); ok && userToken.Valid {
		ctx.Writef("uid:%v exp:%v\n",reflect.TypeOf(claims["uid"]),reflect.TypeOf(claims["exp"]))
		ctx.Writef("uid:%s\n",claims["uid"].(string))

		switch exp := claims["exp"].(type) {
		case float64:
			ctx.Writef("exp at :%f\n",exp)
			ctx.Writef("exp at :%s\n" ,time.Unix(int64(exp),0))
		case json.Number:
			v, _ := exp.Int64()
			ctx.Writef("exp at :%v\n",v)
			ctx.Writef("exp at :%s\n" ,time.Unix(v,0))
		}

	}
	ctx.Writef("This is an authenticated request\n")
	ctx.Writef("Claim content:\n")
	ctx.Writef("%s", userToken.Signature)
	//TODO
}

func Todo(ctx iris.Context){
	log.Println("Todo:call")
	//TODO
}

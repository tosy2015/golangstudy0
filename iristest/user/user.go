package user

import (
	"github.com/kataras/iris"
	"log"
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
	//TODO
}

func Todo(ctx iris.Context){
	log.Println("Todo:call")
	//TODO
}

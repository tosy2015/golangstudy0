package real

import (
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
	log.Println("getList:call")
	//TODO
}

func Todo(ctx iris.Context){
	log.Println("Todo:call")
	//TODO
}

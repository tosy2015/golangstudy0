package main

import (
	"context"
	"fmt"
	"github.com/golangstudy0/iristest/jwtconfig"
	"github.com/golangstudy0/iristest/real"
	"github.com/golangstudy0/iristest/user"
	"github.com/golangstudy0/server/arith/arith"
	pb "github.com/golangstudy0/server/hello/hello"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	irisRecover "github.com/kataras/iris/middleware/recover"
	"google.golang.org/grpc"
	"log"
	"net/rpc"
	"os"
	"time"
)

const (
	addressRpc  = "localhost:1234"
	addressGrpc = "localhost:2234"
	defaultName = "tosy"
)

//docker run  --rm -it -v /Users/tosy/go:/go golang:stretch env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags -s -a -installsuffix cgo -o /go/src/github.com/golangstudy0/iristest/web.out  /go/src/github.com/golangstudy0/iristest/main.go
//docker build -t stretch-web:latest .
//docker run -itd -p 8080:8080 stretch-web:latest
//docker stop id( docker ps -a)
//docker rm id( docker ps -a)

func main() {
	app := iris.New()
	//跨域
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowCredentials: true,
	})
	//log
	app.Logger().SetLevel("debug")
	app.Use(logger.New())
	//recoer
	app.Use(irisRecover.New())
	//jwt
	j := jwtconfig.Get()

	app.Handle("GET", "/", func(ctx iris.Context) {
		ctx.HTML("<h1>Welcome</h1>")
	})

	v1 := app.Party("/api/v1",crs).AllowMethods(iris.MethodOptions)
	{
		r := v1.Party("/real")
		r.Use(j.Serve)
		r.Get("/getList",real.New().GetList)

		u := v1.Party("/user")
		u.Get("/login",user.New().Login)
	}

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	//rpc demo
	app.Get("/rpc", func(ctx iris.Context) {
		client, err := rpc.DialHTTP("tcp", addressRpc)
		if err != nil {
			//log.Fatal("dialing:", err)
			fmt.Println("xxx dialing",err)
		}
		args := &arith.Args{7, 8}
		var reply int
		err = client.Call("Arith.Multiply", args, &reply)
		if err != nil {
			//log.Fatal("arith error:", err)
			log.Println("arith error:", err)
		}
		fmt.Printf("Arith: %d*%d=%d", args.A, args.B, reply)
		ctx.JSON(iris.Map{"args.A": args.A, "args.B": args.B, "reply": reply})
	})
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	//gprc demo
	app.Get("/grpc", func(ctxi iris.Context) {
		defer func() {
			if e := recover(); e != nil {
				fmt.Printf("Panicing %s\r\n", e)
			}
		}()
		conn, err := grpc.Dial(addressGrpc, grpc.WithInsecure())
		if err != nil {
			log.Print("did not connect: %v", err)
		}
		defer conn.Close()
		c := pb.NewGreeterClient(conn)

		// Contact the server and print out its response.
		name := defaultName
		if len(os.Args) > 1 {
			name = os.Args[1]
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
		if err != nil {
			log.Printf("could not greet: %v", err)
		}
		ctxi.JSON(iris.Map{"Greeting": r.Message})
	})

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	log.Println("start...")
	err := app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
	if err != nil {
		log.Println("err",err)
	}
}

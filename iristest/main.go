package main

import (
	"context"
	"fmt"
	"log"
	"net/rpc"
	"os"
	"time"

	"github.com/golangstudy0/server/arith/arith"
	"github.com/kataras/iris"
	"google.golang.org/grpc"

	"github.com/kataras/iris/middleware/logger"
	irisRecover "github.com/kataras/iris/middleware/recover"

	pb "github.com/golangstudy0/server/hello/hello"
)

const (
	addressRpc  = "localhost:1234"
	addressGrpc     = "localhost:2234"
	defaultName = "tosy"
)

//docker run  --rm -it -v /Users/tosy/go:/go golang:stretch env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags -s -a -installsuffix cgo -o /go/src/github.com/golangstudy0/iristest/web.out  /go/src/github.com/golangstudy0/iristest/main.go
//docker build -t stretch-web:latest .
//docker run -itd -p 8080:8080 stretch-web:latest
//docker stop id( docker ps -a)
//docker rm id( docker ps -a)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	// Optionally, add two built'n handlers
	// that can recover from any http-relative panics
	// and log the requests to the terminal.
	app.Use(irisRecover.New())
	app.Use(logger.New())

	// Method:   GET
	// Resource: http://localhost:8080
	app.Handle("GET", "/", func(ctx iris.Context) {
		ctx.HTML("<h1>Welcome</h1>")
	})

	// same as app.Handle("GET", "/ping", [...])
	// Method:   GET
	// Resource: http://localhost:8080/ping
	app.Get("/ping", func(ctx iris.Context) {
		ctx.WriteString("pong")
	})

	// Method:   GET
	// Resource: http://localhost:8080/hello
	app.Get("/hello", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "Hello Iris!"})
	})

	app.Get("/rpc", func(ctx iris.Context) {
		//defer func() {
		//	if e := recover(); e != nil {
		//		fmt.Printf("Panicing %s\r\n", e)
		//	}
		//}()
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

	app.Get("/grpc", func(ctxi iris.Context) {
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
	// http://localhost:8080
	// http://localhost:8080/ping
	// http://localhost:8080/hello
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}

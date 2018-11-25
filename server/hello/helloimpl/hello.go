package helloimpl

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"time"

	pb "github.com/golangstudy0/server/hello/hello"
)

// server is used to implement hello.GreeterServer.
type Hello struct{}

var (
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
)

// SayHello implements hello.GreeterServer
func (s *Hello) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	set, err := client.SetNX("key", "value", 5*time.Second).Result()
	fmt.Println(set, err)

	if !set {
		return &pb.HelloReply{Message:"wait..." + in.Name},nil
	}

	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

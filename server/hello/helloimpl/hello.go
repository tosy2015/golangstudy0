package helloimpl

import (
	"context"

	pb "github.com/golangstudy0/server/hello/hello"
)

// server is used to implement hello.GreeterServer.
type Hello struct{}

// SayHello implements hello.GreeterServer
func (s *Hello) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

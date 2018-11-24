package main

import (
	"log"
	"net"

	pb "github.com/golangstudy0/server/hello/hello"
	"github.com/golangstudy0/server/hello/helloimpl"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

//protoc -I hello/ hello/hello.proto --go_out=plugins=grpc:hello
const (
	port = ":2234"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &helloimpl.Hello{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

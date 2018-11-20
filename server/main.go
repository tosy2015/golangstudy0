package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"

	"github.com/golangstudy0/server/srvrpc"
)

func main() {
	arith := new(srvrpc.Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	// go http.Serve(l, nil)
	http.Serve(l, nil)
}

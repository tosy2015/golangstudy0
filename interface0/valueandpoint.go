package main

import (
	"github.com/golangstudy0/interface0/reciever"
)

func main() {
	var r reciever.Reciverdefault
	r.Recieve()

	rr := new(reciever.Reciverdefault)
	rr.Recieve()
}

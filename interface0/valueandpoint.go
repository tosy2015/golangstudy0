package main

import (
	"fmt"

	"github.com/golangstudy0/interface0/reciever"
)

func main() {
	var r reciever.Reciverdefault
	r.Recieve()

	rr := new(reciever.Reciverdefault)
	rr.Recieve()

	rrr := reciever.Reciverdefault{Age: 1, Sex: 2}
	// rrr.Recieve()
	test(&rrr)

	fmt.Println(r.Age, r.Sex)
	fmt.Println(rr.Age, rr.Sex)
	fmt.Println(rrr.Age, rrr.Sex)

	reciever.GetInvisible(rrr)

}

func test(n reciever.Reciever) {
	n.Recieve()
}

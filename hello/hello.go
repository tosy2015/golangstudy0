package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type world interface {
	Hi()
}

type hello interface {
	Hi()
}

type test struct {
	a	int
}
func (*test) Hi(){
	fmt.Println("hhe")
}

//var _ hello = (*test)(nil)

func main() {
	t1 := test{a : 1}
	//t1.Hi()
	t2 := (world)(&t1)
	if hl , ok := t2.(hello); ok {
		hl.Hi()
	}
	fmt.Println("hello")

	r, e := http.Get("https://www.tosylab.com")
	if e != nil {
		fmt.Println(e)
	}

	io.Copy(os.Stdout, r.Body)
	if err := r.Body.Close(); err != nil {
		fmt.Println(err)
	}
}

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
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

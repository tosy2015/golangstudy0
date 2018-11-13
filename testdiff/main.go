package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("begin test")

	arr := [1024 * 1024][8]int{}
	for i := 0; i < 8; i++ {
		for j := 0; j < 1024*1024; j++ {
			arr[j][i] = 1
		}
	}
	//横向
	sum := 0
	t1 := time.Now() // get current time
	for i := 0; i < 8; i++ {
		for j := 0; j < 1024*1024; j++ {
			sum += arr[j][i]
		}
	}
	elapsed := time.Since(t1)
	fmt.Println("x ", elapsed)

	//纵向
	sum = 0
	t2 := time.Now() // get current time
	for j := 0; j < 1024*1024; j++ {
		for i := 0; i < 8; i++ {
			sum += arr[j][i]
		}
	}
	elapsed = time.Since(t2)
	fmt.Println("x ", elapsed)

}

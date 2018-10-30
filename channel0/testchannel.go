package main

import (
	"log"
	"time"
)

// Result contains the result of a search.
type Result struct {
	Name string
}

func main() {
	results := make(chan *Result)
	//this make deadlock
	// go func() {
	// 	time.Sleep(5 * time.Second)
	// 	close(results)
	// }()
	// log.Println("in for ...")
	for result := range results {
		log.Printf("1" + result.Name)
	}
	log.Println("begin 10 sec sleep")
	time.Sleep(10 * time.Second)
}

package reciever

import (
	"log"
)

//Reciever for test
type Reciever interface {
	Recieve()
}

//Reciverdefault 1
type Reciverdefault struct {
	Age  int
	name string
	Sex  int
}

// //Recieve 1		冲突
// func (m Reciverdefault) Recieve() {
// 	log.Println("m")
// }

//Recieve 2
func (m *Reciverdefault) Recieve() {
	log.Println("m*")
	m.Age = 0
}

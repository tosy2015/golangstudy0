package reciever

import (
	"log"
)

//Reciever for test
type Reciever interface {
	Recieve()
}

//Reciverdefault 1
type Reciverdefault struct{}

// //Recieve 1		冲突
// func (m Reciverdefault) Recieve() {
// 	log.Println("m")
// }

//Recieve 2
func (m *Reciverdefault) Recieve() {
	log.Println("m*")
}

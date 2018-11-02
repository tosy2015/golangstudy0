package reciever

import (
	"fmt"
)

//GetInvisible for test
func GetInvisible(m Reciverdefault) {
	m.name = "haha"
	fmt.Println(m.name)
}

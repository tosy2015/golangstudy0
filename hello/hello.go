package main

import (
	"fmt"
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

type test2 struct {
	a	int
}

func (*test) Hi(){
	fmt.Println("1111")
}

func (*test2) Hi(){
	fmt.Println("2222")
}

var _ hello = (*test)(nil)		//编译器校验  nil指针转结构体指针--赋值--interface。校验结构体是否实现了interface

func main() {
	tA := test{a:1}
	iA := (hello)(&tA)
	tC := iA.(*test2)			//panic		指针的结构体name不同于转换后的结构体name，接口实现可能不同！
	tC.Hi()

	t1 := test{a : 1}
	//t1.Hi()
	t2 := (world)(&t1)
	if hl , ok := t2.(hello); ok {
		hl.Hi()
	}
	fmt.Println("hello")

	//r, e := http.Get("https://www.tosylab.com")
	//if e != nil {
	//	fmt.Println(e)
	//}
	//
	//io.Copy(os.Stdout, r.Body)
	//if err := r.Body.Close(); err != nil {
	//	fmt.Println(err)
	//}
}

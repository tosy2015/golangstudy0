package main

import (
	"fmt"
	"reflect"
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
	//数组
	//arrA := [5]int {1,2,3,4,5}
	arrA := [...]int {1,2,3,4,5}
	parrB := &arrA
	copyA := arrA
	fmt.Printf("arr P %p %p %p  \n\n",&arrA,parrB,&copyA )

	fmt.Printf("arr P %p %p %p  \n\n",&(arrA[0]),&(parrB[0]),&copyA[0])


	tA := test{a:1}
	iA := (hello)(&tA)
	//tC := iA.(*test2)			//panic		指针的结构体name不同于转换后的结构体name，接口实现可能不同！
	//tC.Hi()

	//如何判断interface的结构体名呢？
	fmt.Println("haha TypeOf ",reflect.TypeOf(iA))

	v := reflect.ValueOf(iA)				//获取 iA变量 真实值信息，真实Type，Kind		interface的Type是*main.test   Kind是ptr
	fmt.Println("value:", v)
	fmt.Println("type:", v.Type())
	fmt.Println("kind:", v.Kind())

	v2 := reflect.ValueOf(tA)				//tA的真实值信息，	Kind是struct，Type是main.test
	fmt.Println("value:", v2)
	fmt.Println("type:", v2.Type())
	fmt.Println("kind:", v2.Kind())


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

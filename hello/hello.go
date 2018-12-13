package main

import (
	"fmt"
	"reflect"
	"regexp"
	"time"
)

type world interface {
	Hi()
}

type hello interface {
	Hi()
}

type XXX interface {
	hello
}

type YYY struct{
	test2
}
func (*YYY) Hi(){
	fmt.Println("3333")
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

func DumpMethods(Foo interface{}){
	fooType := reflect.TypeOf(Foo)
	for i := 0; i < fooType.NumMethod(); i++ {
		method := fooType.Method(i)
		fmt.Println(method.Name)
	}
}


func main() {
	testch := make(chan struct{})		//不传最大数量，默认是1
	go func() {
		for {
			select{
			case  _,ok:= <-testch:
				fmt.Println("d := <-testch",ok)
				if !ok {
					return
				}
			}
			fmt.Println("for ...overtestchn")
		}
		fmt.Println("overtestchn")
	}()

	//testch <- struct{}{}
	testch <- struct{}{}
	close(testch)


	s := "AA[:10_x1111好]sdf[:10_11哈哈11]kjjk[:10_11哈哈]jkfBB[:10_哈哈11]zzz[:1_哈哈]kAAkjjk[:12]jBB"
	r := regexp.MustCompile(`\[:[^\]]+\]`)

	//testarr := make([]string,0,10)
	//testarr = append(testarr, "11")
	//testarr = append(testarr, "11")
	//fmt.Println(testarr)

	tN := time.Now()
	for i:=0 ; i < 100000 ; i++{
		r.FindAllStringSubmatch(s,-1)
		//result :=
		//fmt.Println(result)
	}
	elapsed := time.Since(tN)
	fmt.Println("App elapsed: ", elapsed)



	//var c chan struct{}	//c = nil
	//go func() {
	//	<- c				//死锁
	//}()
	//c <- struct {}{}		//死锁							all goroutines are asleep - deadlock!

	//无缓冲
	//c := make(chan struct{})
	//go func() {
	//	<- c
	//}()
	//c <- struct {}{}				//不go取出，主线程死锁

	//有缓冲
	//c := make(chan struct{},1)
	//c <- struct {}{}				//正常
	//close(c)
	//c <- struct {}{}				//close之后发送，panic

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

	fmt.Println(".........")
	DumpMethods(new(YYY))
	iXX := iA.(XXX)
	fmt.Println(".........")
	DumpMethods(iXX)

	y := YYY{}
	y.Hi()
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

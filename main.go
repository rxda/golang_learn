package main

import "fmt"

//
//type Writer interface {
//	Write()
//}
//
//type Author struct {
//	name string
//	Writer
//}
//
//// 定义新结构体，重点是实现接口方法Write()
//type Other struct {
//	i int
//}
//
////func (a Author) Write() {
////   fmt.Println(a.name, "  Write.")
////}
//
//// 新结构体Other实现接口方法Write()，也就可以初始化时赋值给Writer 接口
//func (o Other) Write() {
//	fmt.Println(" Other Write.")
//}
//
//func main() {
//
//	//  方法一：Other{99}作为Writer 接口赋值
//	Ao := Author{"Other", Other{99}}
//	Ao.Write()
//
//	// 方法二：简易做法，对接口使用零值，可以完成初始化
//	//Au := Author{name: "Hawking"}
//	//Au.Write()¡
//}
//func init()  {
//	fmt.Println("a")
//}

type student struct {
	Name string
	Age  int
}

var c1 = make(chan int, 1)
var c2 = make(chan string, 1)

func main() {

	c1 <- 1
	c2 <- "hello"

	select {

	case v1 := <-c1:
		fmt.Println(v1)

	case v2 := <-c2:
		panic(v2)
	}
}

package main

import (
	"fmt"
)

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

const (
	// iota=0 const=1+0=1 iota=0+1=1
	first = 1 + iota

	// iota=1 const=1+1=2 iota=1+1=2
	second

	// iota=2 const=2+2=4 iota=2+1=3
	third = 2 + iota

	// iota=3 const=2+3=5 iota=3+1=4
	forth

	// iota=4 const=2*4=8 iota=4+1=5
	fifth = 2 * iota

	// iota=5 const=2*5=10 iota=5+1=6
	sixth

	// iota=6 const=6 iota=6+1=7
	seventh = iota
)

const (
	a = iota
	b
	c
	d = iota *2
	e
	f
)
// 1 2 4 5 8 10 6
func main() {
	fmt.Println(a,b,c,d,e,f)
}

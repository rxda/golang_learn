package main

import (
	"fmt"
	"time"
)

//select
func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	//c := make(chan int)
	//quit := make(chan int)
	//go func() {
	//	for i := 0; i < 10; i++ {
	//		fmt.Println(<-c)
	//	}
	//	quit <- 0
	//}()
	//fibonacci(c, quit)
	returningo()
}


func returningo(){
	go func() {
		time.Sleep(3)
		fmt.Println("child")
		return
	}()
	time.Sleep(5)
	fmt.Println("main")
	fmt.Println("main")
	fmt.Println("main")
}


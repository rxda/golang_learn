package problem

import "fmt"

// 两个携程，交替打印1-200的自然数
func AlternatePrint() bool {
	c1 := make(chan struct{})
	c2 := make(chan struct{})
	done := make(chan struct{})

	go func() {
		for i := 1; i <= 200; i = i + 2 {
			<-c2
			fmt.Println("c1", i)
			c1 <- struct{}{}
		}
	}()
	go func() {
		for i := 2; i <= 200; i = i + 2 {
			<-c1
			fmt.Println("c2", i)
			c2 <- struct{}{}
		}
		done<- struct {}{}
	}()
	c2 <- struct{}{}
	<-done
	return true
}

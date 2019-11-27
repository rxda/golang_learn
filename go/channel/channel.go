package channel

import "fmt"

func twoChannel() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for {
			ch1 <- 1
		}
	}()

	go func() {
		for {
			ch2 <- <-ch1
		}
	}()

	go func() {
		for {
			fmt.Println(<-ch2)
		}
	}()

	select {}
}

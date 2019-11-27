package recover

import (
	"fmt"
	"time"
)

func running() {

	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()
		for {
			time.Sleep(5 * time.Second)
			panic("hello panic")
		}
	}()
	time.Sleep(20 * time.Second)
}

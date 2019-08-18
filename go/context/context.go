package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)


func coroutine(ctx context.Context, id int, wg *sync.WaitGroup) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("goroutine %d finish\n", id)
			wg.Done()
			return
		default:
			fmt.Printf("message from goroutine %d\n", id)
			time.Sleep(time.Second)
		}
	}
}

func main() {
	wg := &sync.WaitGroup{}
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	//ctx, cancel := context.WithCancel(context.Background())
	//go func() {
	//	time.Sleep(10* time.Second)
	//	cancel()
	//}()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go coroutine(ctx, i, wg)
	}

	wg.Wait()
}

func withCancel(){
	// gen generates integers in a separate goroutine and
	// sends them to the returned channel.
	// The callers of gen need to cancel the context once
	// they are done consuming generated integers not to leak
	// the internal goroutine started by gen.
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					return // returning not to leak the goroutine
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancel when we are finished consuming integers

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}
package mysync

import (
	"sync"
)

type myInt struct {
	Val int
	sync.Mutex
}

func testMutex() int{
	my := myInt{}
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(w *sync.WaitGroup) {
			my.Lock()
			my.Val++
			my.Unlock()
			w.Done()
		}(&wg)
	}
	wg.Wait()
	return my.Val
}

//Error Increment ,data race
func ErrorIncrement() int {
	x := 0
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(w *sync.WaitGroup) {
			x++
			w.Done()
		}(&wg)
	}
	wg.Wait()
	return x
}

func CorrectIncrement() int{
	x := 0
	m := sync.Mutex{}
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(w *sync.WaitGroup) {
			m.Lock()
			x++
			m.Unlock()
			w.Done()
		}(&wg)
	}
	wg.Wait()
	return x
}
package problem

import (
	"strings"
	"sync"
)

//测试strings.Builder是不是协程安全的
//结果：不是
func builderTest() int {
	sb := strings.Builder{}
	//done := make(chan struct{})
	wg := sync.WaitGroup{}
	for i:=0;i<1000;i++{
		wg.Add(1)
		go func() {
			sb.WriteString("A")
			wg.Done()
		}()
	}
	wg.Wait()
	return len(sb.String())
}

type safeBuilder struct {
	strings.Builder
	sync.RWMutex
}

func safeBuilderTest() int {
	sb := safeBuilder{}
	//done := make(chan struct{})
	wg := sync.WaitGroup{}
	for i:=0;i<1000;i++{
		wg.Add(1)
		go func() {
			sb.Lock()
			_, _ = sb.WriteString("A")
			sb.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	return len(sb.String())
}
package mysync

import (
	"sync"
)

func testOnce() int{
	one := sync.Once{}
	sum := 0
	for i := 0; i < 10; i++ {
		one.Do(func(){
			sum++
		})
	}
	return sum
}

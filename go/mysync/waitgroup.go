package mysync

import "sync"

func errorWg() bool{
	m := sync.Map{}
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		//waitgroup with value copy, panic: sync: negative WaitGroup counter
		go func(goi int, w sync.WaitGroup) {
			w.Done()
			m.Store(goi, 1)
		}(i, wg)
	}
	wg.Wait()
	for i := 0; i < 1000; i++ {
		if val, ok := m.Load(i); !ok || val != 1 {
			return false
		}
	}
	return true
}


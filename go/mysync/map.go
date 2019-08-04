package mysync

import "sync"

func ErrCoroutineMap() bool {
	m := make(map[int]int, 1000)
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(goi int, w *sync.WaitGroup) {
			m[goi]++
		}(i, &wg)
	}
	wg.Wait()
	for i := 0; i < 1000; i++ {
		if m[i] != 1 {
			return false
		}
	}
	return true
}

func CorrectCoroutineMap() bool {
	m := sync.Map{}
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(goi int, w *sync.WaitGroup) {
			w.Done()
			m.Store(goi, 1)
		}(i, &wg)
	}
	wg.Wait()
	for i := 0; i < 1000; i++ {
		if val, ok := m.Load(i); ok && val == 1 {
			return false
		}
	}

	//other function
	m.LoadOrStore(1001, 1)
	m.Delete(1)
	m.Range(func(key, value interface{}) bool {
		m.Store(key, value.(int)+10)
		return true
	})
	return true
}

func OtherCoroutineMapFunc() (delete bool, v100, v1001 int) {
	m := sync.Map{}
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(goi int, w *sync.WaitGroup) {
			w.Done()
			m.Store(goi, 1)
		}(i, &wg)
	}
	wg.Wait()
	//other function
	m.LoadOrStore(1001, 11)
	m.Delete(1)
	m.Range(func(key, value interface{}) bool {
		m.Store(key, value.(int)+10)
		return true
	})
	_, ok := m.Load(1)
	val, _ := m.Load(100)
	val1001, _ := m.Load(1001)
	return !ok, val.(int), val1001.(int)
}

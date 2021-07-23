package goroutine

import (
	"fmt"
	"sync"
)

//Lock 加锁防止资源竞争
func Lock() {
	mutexLock()
}

func mutexLock() {

	//var lock sync.Mutex
	var x int64
	var wg sync.WaitGroup
	var rwLock sync.RWMutex

	wg.Add(2)
	for i := 0; i < 2; i++ {
		go func() {
			for i := 0; i < 1000; i++ {
				//lock.Lock()
				rwLock.Lock()
				x = x + 1
				rwLock.Unlock()
				//lock.Unlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(x)
}

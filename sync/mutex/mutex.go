package mutex

import (
	"fmt"
	"sync"
	"time"
)

func MutexAdd(num int) int {
	wg := sync.WaitGroup{}
	mutex := &sync.Mutex{}
	var ss int
	start := time.Now()
	for i := 0; i < num; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mutex.Lock()
			ss += 1
			mutex.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(start))
	return ss
}

package atomic

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func CompareAndSwapUint32(num, oldNum, newNum uint32) bool {
	return atomic.CompareAndSwapUint32(&num, oldNum, newNum)
}

func AtomicAdd(total int) int32 {
	var num int32
	var wg sync.WaitGroup
	start := time.Now()
	for i := 0; i < total; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt32(&num, 1)
		}()
	}
	wg.Wait()
	fmt.Println("time:", time.Now().Sub(start))
	return num
}
func Add(total int) int32 {
	var num int32
	//var wg sync.WaitGroup
	start := time.Now()
	for i := 0; i < total; i++ {
		//wg.Add(1)
		//go func() {
		//defer wg.Done()
		//atomic.AddInt32(&num, 1)
		//}()

		num += 1
	}
	//wg.Wait()
	fmt.Println("time:", time.Now().Sub(start))
	return num
}

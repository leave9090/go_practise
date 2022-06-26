package atomic

import (
	"sync/atomic"
)

func CompareAndSwapUint32(num, oldNum, newNum uint32) bool {
	return atomic.CompareAndSwapUint32(&num, oldNum, newNum)
}

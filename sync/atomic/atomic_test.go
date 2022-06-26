package atomic

import (
	"fmt"
	"testing"
)

func TestCompareAndSwapUint32(t *testing.T) {
	var num, oldNum, newNum uint32
	oldNum = 1
	newNum = 2
	fmt.Println(CompareAndSwapUint32(num, oldNum, newNum))
}

package atomic

import (
	"fmt"
	"testing"
)

func TestCompareAndSwapUint32(t *testing.T) {
	var num, oldNum, newNum uint32
	num = 1
	oldNum = 1
	newNum = 2
	fmt.Println(CompareAndSwapUint32(num, oldNum, newNum))
}

func TestAtomicAdd(t *testing.T) {
	fmt.Println(AtomicAdd(9999999))
}

func TestAdd(t *testing.T) {
	fmt.Println(Add(9999999))

}

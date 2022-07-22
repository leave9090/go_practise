package mem

import (
	"fmt"
	"unsafe"
)

func Men() {
	max := make([]int, 0)
	for i := 0; i < 100000; i++ {
		max = append(max, i)
	}
	fmt.Println(max)
	fmt.Println(unsafe.Sizeof(max) + unsafe.Sizeof([100000]int{}))
}

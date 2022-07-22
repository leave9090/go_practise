package routine

import "fmt"

func Routine() {
	fmt.Println("routine")
	go Routine()
}

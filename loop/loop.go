package loop

import "fmt"

func SumLoop() int {
	sum := 1
	for sum < 10 {
		sum += sum
	}
	return sum
}

func LoopSlice() {
	s := []int{1, 2, 3, 4, 5}
	for i, v := range s {
		fmt.Println("index:", i, "value:", v)
	}
}

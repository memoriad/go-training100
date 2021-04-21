package list

import "fmt"

func PrintArray() {
	var arr = [5]int{1, 2, 3, 4, 5}

	fmt.Println(arr[3])
}

func PrintSlice() {
	s := []int{3, 5, 676, 3, 4}
	fmt.Println(&s[0])
	s = append(s, 100, 1000, 10000)
	fmt.Println(&s[0])
}

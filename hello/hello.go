package hello

import "fmt"

func Hello(text string) {
	fmt.Println(text)
}

func Sum(a, b int) int {
	return a + b
}

func Swap(a int, b int) (int, int) {
	return b, a
}
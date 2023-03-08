package main

import "fmt"

func main() {
	a := []int{1, 5, 6, 6}

	fmt.Printf("This is result: %v\n", Grow(a))
}

func Grow(arr []int) int {
	var res int = 1

	for i := 0; i < len(arr); i++ {
		res *= arr[i]
	}
	return res
}

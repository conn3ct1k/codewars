package main

import "fmt"

func main() {
	output := reversedSeq(5)

	fmt.Printf("This is output: %v\n", output)
}

func reversedSeq(n int) []int {
	result := []int{}

	for i := n; i >= 1; i-- {
		result = append(result, i)
	}
	return result
}

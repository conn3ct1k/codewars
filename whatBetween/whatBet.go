package main

import "fmt"

func Between(a, b int) []int {
	var result []int // initialize the count variable to 0

	for i := a; i <= b; i++ { // loop through the array
		result = append(result, i) // if the current value is true
	}
	return result
}

func main() {
	a := 1

	b := 10

	fmt.Printf("here is result: %v\n", Between(a, b))
}

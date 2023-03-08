package main

import "fmt"

func main() {
	a := []int{2, 4, 5, 5}

	double := Mult(a)
	fmt.Println(double)
}

func Mult(x []int) []int {
	double := make([]int, len(x))
	for i, v := range x {
		double[i] = v * 2
	}
	return double
}

/*
	in go, arrays and slices are data structure that consists of an ordered of elements.
	These data collections are great to use when you want to work with many related values.
	They enable you to keep data together that belongs together, condense your code, and perform
	the same methods and operations on multiple values at once.

	Although arrays and slices in go both ordered sequences of elements, there are significant
	differences between the two. An array in go a data structure that consists of an ordered
	sequence of elements that has its capacity defined at creation time. Once an array has allocated
	its size, the size can no longer be changed. A slice, of the other hand, is variable length
	version of an array, providing more flexibility for developers using these data structures.
	Slices constitute what you would think of as arrays in other languages.

	Given these differences, there are specific situations when would use one over the other. if
	you are new to Go, determining when to use them can be confusing: Although the versatility
	of slices make them a more appropriate choice in most situations, there are specific instances
	in which arrays can optimize the performance of your program.

*/

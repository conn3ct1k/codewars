package main

import "fmt"

func CountTrueValues(arr []bool) int {
	count := 0 // initialize the count variable to 0

	for _, value := range arr { // Loop through the array
		if value == false { // If the current value is true
			count++ // Increment the count variable
		}
	}
	return count
}

func main() {
	arr := []bool{true, false, true, true, false}
	count := CountTrueValues(arr)
	fmt.Println(count)
}

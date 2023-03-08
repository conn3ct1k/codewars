package main

import (
	"fmt"
	"strconv"
)

func main() {
	input := "5 3 6 4 2"

	output := HighAndLow(input)

	fmt.Println(output)
}

// func HighAndLow(in string) string {
// 	// Split the input string into a slice of strings
// 	numbers := strings.Split(in, " ")

// 	// initialize variables to hold the min and max values
// 	min, max := 0, 0

// 	// Loop through the number and update minx and max
// 	for _, num := range numbers {
// 		// Convert the string to an integer

// 		n, _ := strconv.Atoi(num)
// 		// Update min if n is smaller than the current value of min, or if min is still 0(which means it hasn't been updated yet)
// 		if n < min || min == 0 {
// 			min = n
// 		}

// 		// Update max if n is arger than the current value of max
// 		if n > max {
// 			max = n
// 		}

// 	}
// 	// return the min and max values as a string
// 	return fmt.Sprintf("%d %d", max, min)
// }

func HighAndLow(in string) string {
	// initialize variable to hold the min and max values
	var min, max int
	var num string

	// initialize a slice to hold the individual numbers

	numbers := make([]int, 0)

	// loop through each character in the input string

	for _, c := range in {
		// if the current character is a space, parse the current number and add it to the numbers slice
		if c == ' ' {
			n, _ := strconv.Atoi(num)
			numbers = append(numbers, n)
			num = ""
		} else {
			num += string(c)
		}
	}

	// Parse the final number and add it to the numbers slice

	n, _ := strconv.Atoi(num)
	numbers = append(numbers, n)

	// Find the minimum and maximum values it  the number slice

	for i, n := range numbers {
		if i == 0 {
			min, max = n, n
		} else {
			if n < min {
				min = n
			}
			if n > max {
				max = n
			}
		}
	}
	return fmt.Sprintf("%d %d", max, min)
}

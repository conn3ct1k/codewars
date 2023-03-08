package main

import "fmt"

func main() {
	month := 6
	quarter := QuarterOf(month)
	fmt.Printf("Month d% is in Quarter %d\n", month, quarter)
}

func QuarterOf(month int) int {
	if month < 1 || month > 12 {
		panic("Invalid month number")
	}
	switch month {
	case 1, 2, 3:
		return 1
	case 4, 5, 6:
		return 2
	case 7, 8, 9:
		return 3
	case 10, 11, 12:
		return 4
	default:
		panic("Invalid month number")

	}
}

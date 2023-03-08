package main

import (
	"fmt"
)

func main() {
	var a int

	fmt.Print("Input your number: ")
	fmt.Scan(&a)
	fmt.Print("Your negative number is:")
	fmt.Println(retNeg(a))
}

func retNeg(n int) int {
	if n != 0 && n > 0 {
		return -n
	}
	return n
}

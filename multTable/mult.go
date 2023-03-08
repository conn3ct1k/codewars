package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Print("Please enter your digit to multiply it: ")
	fmt.Scan(&n)
	fmt.Println(mult(n))
}

func mult(n int) string {
	table := ""

	for i := 1; i <= 10; i++ {

		r := i * n
		line := fmt.Sprintf("%d * %d = %d\n", i, n, r)

		table += line

	}
	return table[:len(table)-1] // this line returns the `result` string, but with the last character removed. 
								// this is because the last character of the string is a newline character that was added after the last
								// 
}

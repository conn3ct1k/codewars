package main

import "fmt"

func main() {
	var str string
	fmt.Print("Please enter word: ")
	fmt.Scan(&str)
	fmt.Println("Your result is:", convertedCase(str))
}

// This function takes in a string 'str' and returns the converted string.

func convertedCase(str string) string {
	converted := "" // here saving converted char

	for _, char := range str { // here loop through each charachter in string
		if char >= 'A' && char <= 'Z' { // here if statement
			converted += string(char + 32)
		} else if char >= 'a' && char <= 'z' {
			converted += string(char - 32)
		} else {
			converted += string(char)
		}
	}
	return converted
}

package main

import "strings"

// Define a function called CapitalizeLetter that takes a string argument and returns a string.
func CapitalizeLetter(str string) string {
	// Initialize an empty string variable called s.
	var s string

	// Set the first letter of the input string to uppercase and add it to the s variable.
	s = strings.ToUpper(string(str[0]))

	// Start a loop that will go through every character in the input string starting from the second character.
	for i := 1; i < len(str); i++ {
		// Add the current character to the s variable.
		s += string(str[i])

		// Check if the current character is a space.
		if str[i] == ' ' {
			// If the current character is a space, set the next character to uppercase and add it to the s variable.
			s += strings.ToUpper(string(str[i+1]))
			// Increment i by 1 to skip the next character since it has already been added to the s variable.
			i++
		}
	}

	// Return the s variable.
	return s
}

// func capitalizeWords(str string) string {
// 	var s string

// 	s = strings.ToUpper(string(str[0]))

// 	for i := 1; i < len(str); i++ {

// 		s += string(str[i])

// 		if str[i] == ' ' {
// 			s += strings.ToUpper(string(str[i+1]))
// 			i++
// 		}

// 	}

// 	return s // your code here...
// }

func main() {
	// Example usage
	str := "the quick brown fox jumps over the lazy dog"
	capitalizedStr := CapitalizeLetter(str)
	println(capitalizedStr)
}

package main

import "fmt"

func main() {
	str := "This is string where we should remove spaces"

	output := removeSpaces(str)

	fmt.Printf("Output: %v\n", output)
}

func removeSpaces(word string) string {
	output := make([]byte, len(word))

	j := 0

	for i := 0; i < len(word); i++ {
		if word[i] != ' ' {
			output[j] = word[i]
			j++
		}
	}

	return string(output[:j])
}

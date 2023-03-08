package main

import (
	"fmt"
	"strings"
)

func main() {
	result := altCap("cosmos and space")
	fmt.Println(result)
}

func altCap(st string) []string {
	evenCap := ""
	oddCap := ""

	for i, v := range st {
		if i%2 == 0 {
			evenCap += strings.ToUpper(string(v))
			oddCap += string(v)
		} else {
			evenCap += string(v)
			oddCap += strings.ToUpper(string(v))
		}
	}
	return []string{evenCap + ",", oddCap}
}

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(countCaracter("Hello World"))
}

func countCaracter(input string) int {
	input = strings.Replace(input, " ", "", -1)
	return len(input)
}
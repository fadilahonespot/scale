package main

import "fmt"

func main() {
	fmt.Println(faktorial(6))
}

func faktorial(number int) int {
	if number == 0 {
		return 1
	}
	return number * faktorial(number-1)
}
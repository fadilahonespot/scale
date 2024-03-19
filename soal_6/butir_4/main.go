package main

import "fmt"

func main() {
	input := []int{4, 2, 3, 1, 7, 8, 5, 12, 9, 10}
	large := largestNumber(input)
	small := smallesNumber(input)
	
	fmt.Println("Bilangan terbesar: ", large)
	fmt.Println("Bilangan terkecil: ", small)
}

func largestNumber(input []int) int {
	var largest int
	for _, v := range input {
		if v > largest {
			largest = v
		}
	}
	return largest
}

func smallesNumber(input []int) int {
	var smallest int
	if len(input) != 0 {
		smallest = input[0]
	}
	for _, v := range input {
		if v < smallest {
			smallest = v
		}
	}
	return smallest
}
package main

import (
	"fmt"
	"math/rand"
)

func printRandomNumbers(numbers []int, remaining []int) {
	if len(numbers) == 10 {
		return
	}

	index := rand.Intn(len(remaining))
	number := remaining[index]

	numbers = append(numbers, number)
	fmt.Printf("%d ", number)

	remaining = append(remaining[:index], remaining[index+1:]...)
	printRandomNumbers(numbers, remaining)
}

func main() {
	fmt.Println("Urutan angka acak:")
	var numbers []int
	remaining := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	printRandomNumbers(numbers, remaining)
	fmt.Println()
}

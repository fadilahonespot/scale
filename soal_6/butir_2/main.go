package main

import "fmt"

func main() {
	fmt.Println(isPalindrome("recer"))
}

func isPalindrome(input string) bool {
	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-1-i] {
			return false
		}
	}
	return true
}
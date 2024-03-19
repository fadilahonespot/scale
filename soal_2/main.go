package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Kata/kalimat yang di input: ")
	scanner.Scan()
	input := scanner.Text()
	isPalindrome := isPalindrome(input)
	if isPalindrome {
		fmt.Println("Merupakan palindrom")
		return
	}

	fmt.Println("Bukan palindrom")
}

func isPalindrome(input string) bool {
	var res string
	for i := len(input); i > 0; i-- {
		res += string(input[i-1])
	}
	if res != input {
		return true
	}
	return false
}

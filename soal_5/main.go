package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Masukkan waktu dalam format 12 jam (HH:MM AM/PM): ")
	scanner.Scan()
	input := scanner.Text()

	convertedTime, err := convertTimeTo24HourFormat(input)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Waktu dalam format 24 jam:", convertedTime)
}

func convertTimeTo24HourFormat(timeStr string) (string, error) {
	if timeStr[2] != ':' {
		return "", fmt.Errorf("invalid input format")
	}

	var hour, minute, second int
	_, err := fmt.Sscanf(timeStr[:2], "%d", &hour)
	if err != nil || hour < 0 || hour > 12 {
		return "", fmt.Errorf("hour tidak valid")
	}

	_, err = fmt.Sscanf(timeStr[3:5], "%d", &minute)
	if err != nil || minute < 0 || minute > 59 {
		return "", fmt.Errorf("Minute tidak valid")
	}

	_, err = fmt.Sscanf(timeStr[6:8], "%d", &second)
	if err != nil || second < 0 || second > 59 {
		return "", fmt.Errorf("Second tidak valid")
	}

	mode := strings.Replace(timeStr[8:], " ", "", -1)
	if mode == "PM" && hour != 12 {
		hour += 12
	} else if mode == "AM" && hour == 12 {
		hour = 0
	}

	return fmt.Sprintf("%02d:%02d:%02d", hour, minute, second), nil
}

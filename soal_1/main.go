package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func main() {
    file, err := os.Open("data.txt")
    if err != nil {
        fmt.Println("Gagal membuka file:", err)
        return
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)

    total := 0
	count := 0
    for scanner.Scan() {
        line := scanner.Text()

        num, err := strconv.Atoi(line)
        if err != nil {
            fmt.Println("Gagal mengonversi angka:", err)
            return
        }

        total += num
		count++
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Error saat membaca file:", err)
        return
    }

	fmt.Println("Total angka pada file:", count)
    fmt.Println("Jumlah semua angka:", total)
}

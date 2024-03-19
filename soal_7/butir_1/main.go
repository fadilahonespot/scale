package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	url := "https://jsonplaceholder.typicode.com/posts"
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(data)
}

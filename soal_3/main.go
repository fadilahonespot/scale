package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Data struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func main() {
	apiUrl := "https://jsonplaceholder.typicode.com/posts"
	response, err := http.Get(apiUrl)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer response.Body.Close()

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var data []Data
	err = json.Unmarshal(responseData, &data) 
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	res, _ := json.Marshal(data)
	fmt.Println("Response from API:", string(res))
}

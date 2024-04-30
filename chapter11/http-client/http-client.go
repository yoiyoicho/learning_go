package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func main() {
	client := http.Client{
		Timeout: 30 * time.Second,
	}

	url := "https://jsonplaceholder.typicode.com/todos/1"
	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodGet,
		url,
		nil)
	if err != nil {
		panic(err)
	}
	res, err := client.Do(req)
	// レスポンスのステータスの数値コードはフィールド StatusCode に、
	// レスポンスコードのテキストはフィールド Status に、
	// レスポンスはio.ReadCloser型のBodyフィールドに格納されている。
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		panic("Unexpected status code")
	}
	fmt.Println(res.Header.Get("Content-Type"))
	var data struct {
		UserID    int    `json:"userId"`
		ID        int    `json:"id"`
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", data)
}

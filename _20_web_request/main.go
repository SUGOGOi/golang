package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		fmt.Println("Error", err)
		return
	}

	defer res.Body.Close()
	// fmt.Println(res)
	// fmt.Println(res)

	data, err1 := io.ReadAll(res.Body)

	if err1 != nil {
		fmt.Println("Error", err1)
		return
	}

	fmt.Println(string(data))
}

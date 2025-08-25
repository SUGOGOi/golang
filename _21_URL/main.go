package main

import (
	"fmt"
	"net/url"
)

func main() {
	myurl := "https://jsonplaceholder.typicode.com/todos/1"

	fmt.Printf("Type %T\n", myurl)

	parseURL, err := url.Parse(myurl)

	if err != nil {
		fmt.Println("Error :", err)
	}

	fmt.Printf("Type %T\n", parseURL)
	fmt.Println("Scheme : ", parseURL.Scheme)
	fmt.Println("Host : ", parseURL.Host)
	fmt.Println("Path : ", parseURL.Path)
	fmt.Println("Raw Query : ", parseURL.RawQuery)
}

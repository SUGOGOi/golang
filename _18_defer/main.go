package main

import "fmt"

func main() {
	fmt.Println("1 ")
	defer fmt.Println("2") //2nd
	defer fmt.Println("4") //1st ====== stack system
	fmt.Println("3")
}

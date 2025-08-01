package main

import "fmt"

func main() {
	helloWorld()

	fmt.Println("1. Add 2 numbers :")
	var a int
	fmt.Println("Enter 1st Number")
	fmt.Scan(&a)

	var b int
	fmt.Println("Enter 2nd Number")
	fmt.Scan(&b)

	var c int = add(a, b)
	fmt.Println("Ans :", c)
}

func helloWorld() {
	fmt.Println("Hello World")
}

func add(a int, b int) int {
	return a + b
}

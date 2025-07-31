package utils

import "fmt"

func PrintMessage(message string) {
	fmt.Println(message)

	printMessage()
}

// In Go, if you declare a function with a lowercase letter at the beginning in a subpackage it's private to that package by default. To make it accessible from outside function name must begin with uppercase letter
func printMessage() {
	fmt.Println("Private Hello World from Utils")
}

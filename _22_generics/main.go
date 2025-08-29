package main

import "fmt"

func printSlice[T int | string | bool](items []T) {
	for index, _ := range items {
		fmt.Println(items[index])

	}
}

// func printStringSlice(items []string) {
// 	for index, _ := range items {
// 		fmt.Println(items[index])

// 	}
// }

func main() {

	printSlice([]int{2, 3, 4, 5})
	printSlice([]string{"sum", "sum", "gogoi"})
	printSlice([]bool{true, false})
}

package main

import "fmt"

func main() {
	fmt.Println("Array")

	// var arr [5]int
	// arr[0] = 1
	// arr[1] = 2
	// arr[2] = 3
	// arr[3] = 4
	// arr[4] = 5

	arr := [5]int{1, 2, 3, 4, 5}

	fmt.Println("Elements :")
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	fmt.Println("Indexes :")
	for i, v := range arr {
		fmt.Println("index :", i, "Value :", v)
	}

	var mstrix [2][3]int = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	fmt.Println(mstrix)

}

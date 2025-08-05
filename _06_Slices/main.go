package main

import "fmt"

func main() {
	// numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// numbers = append(numbers, 11, 12, 13, 14, 15)

	// for i := 0; i < len(numbers); i++ {
	// 	fmt.Println(numbers[i])
	// }

	// fmt.Println(numbers)
	// fmt.Printf("data type : %T\n", numbers)

	// names := []string{}
	// fmt.Println(names)

	numbers := make([]int, 3, 5)
	numbers = append(numbers, 11)
	numbers = append(numbers, 21)
	numbers = append(numbers, 31)

	fmt.Println("Slice :", numbers)
	fmt.Println("Length :", len(numbers))
	fmt.Println("Capacity :", cap(numbers))
}

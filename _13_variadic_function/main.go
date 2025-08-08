package main

import "fmt"

func sum(nums ...int) int { //interface{} = for any type
	total := 0

	for _, num := range nums {
		total = total + num
	}

	return total
}

func main() {

	fmt.Println(sum(1, 2, 3, 4, 411, 313, 2123131))

	// numbers := []int{3, 4, 5, 6}
	// fmt.Println(sum(numbers...))
}

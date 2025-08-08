package main

import "fmt"

func main() {

	// var num int
	// num = 2

	// var ptr *int
	// ptr = &num

	num := "Sumsum"

	ptr := &num
	fmt.Println(ptr)
	fmt.Println(*ptr)

	value := 5
	modifyValue(&value)
	fmt.Println(value)

}

func modifyValue(ptr *int) {
	*ptr = *ptr * 2
}

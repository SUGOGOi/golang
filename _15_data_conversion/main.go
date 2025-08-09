package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int = 1
	fmt.Println("Number value :", num)
	fmt.Printf("Number Type : %T\n", num)

	var data float64 = float64(num)
	fmt.Println("data value :", data)
	fmt.Printf("data Type : %T\n", data)

	// var data2 string = string(num)
	str := strconv.Itoa(num)
	fmt.Println("str value :", str)
	fmt.Printf("str Type : %T\n", str)

	num_string := "123"
	num_int, _ := strconv.Atoi(num_string)
	num_int = num_int + 20
	fmt.Println("num_int value :", num_int)
	fmt.Printf("num_int Type : %T\n", num_int)

	num_string2 := "3.14"
	num_float, _ := strconv.ParseFloat(num_string2, 64)
	fmt.Println("num_float value :", num_float)
	fmt.Printf("num_float Type : %T\n", num_float)
}

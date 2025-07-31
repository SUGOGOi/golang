package main

import "fmt"

func Variables() {
	var a int = 100
	var name string = "Sumsum Gogoi"
	// var b string = 100  ========== Error
	var c = "Hello World"
	var d = 200
	var e float64 = 200.64
	var f bool = true

	const pi = 22 / 7

	person := "Sumsum Gogoi"

	var Public = "public variable" //can access by other files too
	var private = "private variable"

	fmt.Println("a : ", a)
	fmt.Println("name : ", name)
	fmt.Println("c : ", c)
	fmt.Println("d : ", d)
	fmt.Println("e : ", e)
	fmt.Println("f : ", f)
	fmt.Println("pi : ", pi)
	fmt.Println("person : ", person)
	fmt.Println("Public : ", Public)
	fmt.Println("private : ", private)
}

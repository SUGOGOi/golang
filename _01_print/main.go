package main

import "fmt"

func main() {
	age := 25
	name := "Sumsum Gogoi"
	height := 173.21

	fmt.Println("Name :", name, "Age :", age, "Height :", height) // auto space addng between parameters
	fmt.Println("Hello World")                                    // print new line

	fmt.Printf("age is %d", age)
	fmt.Printf("height is %.3f\n", height)
	fmt.Printf("name is %s\n", name)

	fmt.Printf("Type of age is %T\n", age)
	fmt.Printf("Name : %s,  Age : %d,  Height : %.2f.", name, age, height)

}

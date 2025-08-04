package main

import "fmt"

func main() {
	fmt.Println("Error Handling with Underscore")

	ans, err := divide(10, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Divison of 10 by 0 is :", ans)

	// ans, _ := divide(10, 0)
	// fmt.Println("Divison 10 by 0 is :", ans)
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("denominator must not be zero")
	}
	return a / b, nil
}

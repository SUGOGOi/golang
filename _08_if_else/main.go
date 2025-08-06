package main

import "fmt"

func main() {
	age := 18

	if age < 18 {
		fmt.Println("Can't Vote")
	} else if age == 18 {
		fmt.Println("Thinking....")
	} else {
		fmt.Println("can vote")
	}

	y := 10

	if age > 10 && (y < 10 || y == 10) {
		fmt.Println("Hiiiii")
	}

}

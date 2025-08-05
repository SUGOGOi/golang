package main

import "fmt"

func main() {
	// for i := 0; i < 5; i++ {
	// 	fmt.Println(i + 1)
	// }

	//Infinite Loop
	// counter := 0
	// for {
	// 	fmt.Println("Infinite Loop")
	// 	counter++

	// 	if counter == 5 {
	// 		break
	// 	}
	// }

	// range
	// numbers := []int{9, 8, 7, 6, 5}

	// for index, value := range numbers {
	// 	fmt.Println("Index :", index, "Value :", value)
	// }

	name := "Sumsum Gogoi"
	for index, char := range name {
		fmt.Printf("Index : %d, Char : %c \n", index, char)
	}
}

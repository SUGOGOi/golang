package main

import (
	"fmt"
	// "io"
	"os"
)

func main() {
	// file, err := os.Create("example.txt")

	// if err != nil {
	// 	fmt.Println("Error creating file", err)
	// 	return
	// }

	// defer file.Close()

	// content := "Hello"

	// _byte, err2 := io.WriteString(file, content+"\n")

	// fmt.Println(_byte)
	// if err2 != nil {
	// 	fmt.Println("Error writing file", err2)
	// 	return
	// }

	//READ
	// file, err := os.Open("example.txt")
	// if err != nil {
	// 	fmt.Println("Error opening file", err)
	// 	return
	// }

	// defer file.Close()

	// //create buffer

	// buffer := make([]byte, 1024)

	// for {
	// 	n, error_ := file.Read(buffer)

	// 	if error_ == io.EOF {
	// 		break
	// 	}

	// 	if error_ != nil {
	// 		fmt.Println("Error reading file", error_)
	// 		return
	// 	}

	// 	fmt.Println(string(buffer[:n]))
	// }

	content, err := os.ReadFile("example.txt")
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}

	fmt.Println(string(content))

}

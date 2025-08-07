package main

import "fmt"

func main() {

	studentMarks := make(map[string]int)

	studentMarks["Sumsum Gogoi"] = 70
	studentMarks["Manoj"] = 60
	studentMarks["Suraj"] = 50

	// fmt.Println("Sumsum's Marks :", studentMarks["Sumsum Gogoi"])

	// delete(studentMarks, "Sumsum Gogoi")

	// fmt.Println("Sumsum's Marks :", studentMarks["Sumsum Gogoi"])

	//checking if key exist
	// marks, exist := studentMarks["Sumsum Gogoi"]
	// marks2, exist2 := studentMarks["Manoj"]

	// fmt.Println("Sumsum Exist? :", exist)
	// fmt.Println("Sumsum Marks? :", marks)
	// fmt.Println("Manoj Exist? :", exist2)
	// fmt.Println("Manoj Marks? :", marks2)

	for index, value := range studentMarks {

		fmt.Println("Marks of", index, "is :", value)
	}

}

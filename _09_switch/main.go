package main

import "fmt"

func main() {
	// day := 8

	// switch day {
	// case 1:
	// 	fmt.Println("Monday")
	// case 2:
	// 	fmt.Println("Tuesday")
	// case 3:
	// 	fmt.Println("Wednesday")
	// case 4:
	// 	fmt.Println("Thursday")
	// case 5:
	// 	fmt.Println("Friday")
	// case 6:
	// 	fmt.Println("Saturday")
	// case 7:
	// 	fmt.Println("Sunday")
	// default:
	// 	fmt.Println("Holiday")
	// }

	month := "january"

	switch month {
	case "january", "february", "march":
		fmt.Println("Winter")
	case "april", "may", "june":
		fmt.Println("Spring")
	default:
		fmt.Println("Other season")
	}

	temp := 10
	switch { //no expression specified
	case temp < 0:
		fmt.Println("Freezing")
	case temp >= 0 && temp < 10:
		fmt.Println("cold")
	default:
		fmt.Println("warm")
	}

}

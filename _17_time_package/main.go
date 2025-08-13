package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println(currentTime)

	// formatted := currentTime.Format("02-01-2006, 03:04:05")
	formatted := currentTime.Format("02-01-2006, 3:04 PM")
	fmt.Println(formatted)

	layout_str := "02/01/2006" //02-01-2006
	date_str := "25/11/2030"
	formated_str, _ := time.Parse(layout_str, date_str)
	fmt.Println(formated_str)

	//add 1 more day in currentTime
	newTime := currentTime.Add(24 * time.Hour)
	formated_new_Time := newTime.Format("2006/01/02 Monday")
	fmt.Println(formated_new_Time)
}

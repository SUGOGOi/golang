package main

import (
	"fmt"
	"strings"
)

func main() {
	data := "sumsum,gogoi" //csv

	parts := strings.Split(data, ",")

	fmt.Println(parts)
	for index, _ := range parts {
		fmt.Println(parts[index])
	}

	str := "one two three four two two five"

	count := strings.Count(str, "two")
	fmt.Println(count)

	str1 := "     Hello, Go!       "
	str1 = strings.TrimSpace(str1)
	fmt.Println(str1)

	str2 := "Sumsum"
	str1 = "Gogoi"
	combined := strings.Join([]string{str2, str1}, " ")
	fmt.Println(combined)

}

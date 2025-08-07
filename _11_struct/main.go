package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Contact struct {
	Email string
	Phone string
}

type Address struct {
	House int
	Area  string
	State string
}

type Employee struct {
	Person_Details Person
	Person_Contact Contact
	Person_Address Address
}

func main() {

	// var person1 Person

	// person1.Age = 25
	// person1.Name = "Sumsum Gogoi"

	// person2 := Person{
	// 	Name: "Manoj Chetry",
	// 	Age:  27,
	// }

	// // new keyword
	// var person3 = new(Person)
	// person3.Age = 24
	// person3.Name = "Suraj Nepal"
	// fmt.Println("person 1", person1)
	// fmt.Println("person 2", person2)
	// fmt.Println("person 3", person3)

	employee1 := Employee{

		Person_Details: Person{
			Name: "Sumsum Gogoi",
			Age:  25,
		},

		Person_Contact: Contact{
			Email: "sumsumgogoi51@gmail.com",
			Phone: "9876543210",
		},

		Person_Address: Address{
			House: 100,
			State: "Assam",
			Area:  "XYZ",
		},
	}

	fmt.Println(employee1)
}

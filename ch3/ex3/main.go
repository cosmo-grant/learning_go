package main

import "fmt"

func main() {
	type Employee struct {
		firstName string
		lastName  string
		id        int
	}

	e1 := Employee{
		"a",
		"b",
		1,
	}

	e2 := Employee{
		firstName: "c",
		lastName:  "d",
		id:        2,
	}

	var e3 Employee
	e3.firstName = "e"
	e3.lastName = "f"
	e3.id = 3

	fmt.Println(e1)
	fmt.Println(e2)
	fmt.Println(e3)
}

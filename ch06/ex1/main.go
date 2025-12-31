package main

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func MakePerson(firstName string, lastName string, age int) Person {
	return Person{firstName, lastName, age}
}

func MakePersonPointer(firstName string, lastName string, age int) *Person {
	// returning a pointer, so its value escapes
	return &Person{firstName, lastName, age}
}

func main() {
	MakePerson("cosmo", "grant", 33)
	// does not escape; i suppose nothing can escape from main?
	MakePersonPointer("cosmo", "grant", 33)
}

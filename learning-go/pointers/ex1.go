package main

import "fmt"

func main() {
	person := MakePerson("Mike", "Peter", 23)
	personPtr := MakePersonPointer("Mike", "Peter", 23)
	fmt.Println(person, personPtr)
}

type Person struct {
	FirstName, LastName string
	Age int
}

func MakePerson(firstName, lastName string, age int) Person{
	return Person{firstName, lastName, age}
}

func MakePersonPointer(firstName, lastName string, age int) *Person {
	return &Person{firstName, lastName, age}
}

package main


func main() {
	const size = 10000000

	//var people []Person
	people := make([]Person, 0, size+10)

	for i := 0; i < size; i++ {
		people = append(people, Person{"Mike", "Peter", 23})
	}
}

type Person struct {
	FirstName, LastName string
	Age int
}


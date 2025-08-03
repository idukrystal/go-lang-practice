package main

import "fmt"

type Animal interface {
	getSound() string
	getAge() int
}

func (a Dog) getSound() string{
	return "woof !"
}

func (d Dog) getAge() int {
	return d.age
}

type Dog struct {
	name string
	age int
	Animal
}


func print(a Animal) {
	fmt.Println(a.getSound, a.getAge())
}

func main() {
	spark := Dog {
		name: "Bingo",
		age: 5,
	}
	lion := struct {
		Animal
		temprament float64
	}{
		temprrament: 68.5,
	}
	lion.makeSound()
	print(spark.Animal)
}


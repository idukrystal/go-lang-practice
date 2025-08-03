package main

import "fmt"

type myInt int

func (i myInt) double() int {
	fmt.Println("getting value")
	return int(i * 2)
}

func main() {
	var varriable any
	var int  myInt = 6
	varriable = int
	fmt.Println(int, varriable, varriable.(myInt).double())
}

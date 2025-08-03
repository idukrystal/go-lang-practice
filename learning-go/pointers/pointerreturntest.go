package main

import "fmt"

func getStringPointer(value string) *string {
	//val := value
	//ptr := &val
	return &value //ptr
}

func main() {
	s := "cab"
	ptr1 := &s

	ptr2 := getStringPointer("bag")

	ptr3 := func(s string) *string { return &s }("bog")

	ptr4 := new(string)
	*ptr4 = "abs"

	fmt.Println(*ptr1, *ptr2, *ptr3, *ptr4)
}

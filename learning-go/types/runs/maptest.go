package main

import "fmt"

func main() {
	var slice = make([]int, 6, 10)
	slice2 := slice[:]
	fmt.Println(slice, cap(slice), len(slice))
	slice = append(slice, 6)
	_ = append(slice2, 5)
	fmt.Println(slice, cap(slice), len(slice))

}

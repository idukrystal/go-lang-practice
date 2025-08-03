package main

import "fmt"

type Converter[T any, U any] interface {
	Convert(t T) U
}


func ConvertSlice[T any, U any](c Converter[T, U], slice []T) []U {
	newSlice := make([]U, len(slice))
	for i, val := range slice {
		newSlice[i] = c.Convert(val)
	}
	return newSlice
}


type StringToIntConvFunc func(string) int
func (sticf StringToIntConvFunc) Convert(s string) int {
	return sticf(s)
}
func getLength(s string) int  {return len(s)}

type multiplier struct {
	mul int
}
func (m multiplier) Convert(i int) int { 
	return m.mul * i
}


func main() {
	names := []string{ "Mike", "Joy", "Paulina", "James"}

	converter := StringToIntConvFunc(getLength)

	nameLengths := ConvertSlice(converter, names)

	doubler := multiplier{2}

	fmt.Println(ConvertSlice(doubler, nameLengths))
}

package main

import "fmt"

type Counter interface {
	increment(int) int
	getCount() int
}

type Integer int

func (i *Integer) increment(incrementBy int) int {
	 *i += Integer(incrementBy)
	return int(*i)
}

func (i Integer) getCount() int {
	return int(i)
}

func main() {
	var counter Counter

	counter = new(Integer)
	counter.increment(5)
	fmt.Println(counter.getCount())
}

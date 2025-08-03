package main

import "fmt"

type A interface {
	a() B
}

type B interface {
	b() A
}

type ABC interface {
	A
	B
	c() ABC
}

type D struct {
	i int
	s string
}

func (d D) a() B {
	return d
}

func (d D) b() A {
	return d
}

func (d D) c() ABC{
	fmt.Printf("%s: %d\n", d.s, d.i)
	return d
}

func aBCFunc(abc ABC) {
	abc.c()
}

func main() {
	d := D{5, "boo"}
	aBCFunc(d)
}

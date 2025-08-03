package main

import "fmt"

func getPrefixer(prefix string) func(string) string {
	return func(base string) string {
		return prefix + " " + base
	}
}

func main() {
	helloPrefixer := getPrefixer("hello")
	msg := helloPrefixer("james")
	fmt.Println(msg)
	fmt.Println(helloPrefixer("maria"))
}

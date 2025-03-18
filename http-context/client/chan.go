package main

import (
	"fmt"
)

func main() {
	c := make(chan string)

	go func() {
		fmt.Println(<-c)
		close(c)
	}()

	c <- "good nonsense"
	//fmt.Println(<-c)

	for {
		select {
		case <-c:
			fmt.Println("found")
			return
		default:
			print("default\n")
		}
	}
}


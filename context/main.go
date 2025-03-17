package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		var s string
		fmt.Scan(&s)
		cancel()
	}()
	sleepAndTalk(ctx, 5*time.Second, "Hello")

}

func sleepAndTalk(ctx context.Context, d time.Duration, msg string) {
	a := after(d)
L:
	for {
		select {
		case <-a:
			fmt.Println(msg)
			break L
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			break L
		default:
			time.Sleep(2*time.Second)
			fmt.Println("quit: ")
		}
	}
}

func after(d time.Duration) chan bool {
	c := make(chan bool)
	go func() {
		time.Sleep(d)
		c <- true
	}()
	return c
}


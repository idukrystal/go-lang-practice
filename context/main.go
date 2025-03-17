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
			fmt.Println("Hold on...")
		}
	}

}

func after(d time.Duration) chan bool {
	c := make(chan bool)
	go func() {
		time.Sleep(d)
		fmt.Println("sending out now")
		c <- true
		fmt.Println("sent")
	}()
	return c
}

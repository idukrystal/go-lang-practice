package main

import (
	"context"
	"fmt"
	"sync"
)

func main() {
	ctx := context.Background()
	fmt.Println(ctx.Value("foo"))
	ctx = context.WithValue(ctx, "foo", "bar")
	fmt.Println(ctx.Value("foo"))
	ctx = context.WithValue(ctx, 45, 60)
	fmt.Println(ctx.Value("foo"))
    var wg sync.WaitGroup
	wg.Add(1)
	go func(ctx context.Context) {
		ctx, _  = context.WithCancel(ctx)
		fmt.Println(ctx.Value("foo"), "background")
		wg.Done()
	}(ctx)
	wg.Wait()
}

package main

import (
	"context"
	"fmt"
	"github.com/idukrystal/go-lang-practice/http-context/log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", log.Decorator(handler))
	panic(http.ListenAndServe("127.0.0.1:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx = context.WithValue(ctx, int(42), int64(100))
	log.Println(ctx, "Handler Started")
	defer log.Println(ctx, "Handler Ended")

	fmt.Printf("Value For: %s\n", ctx.Value("foo"))
	select {
	case <- time.After(5 * time.Second):
		fmt.Fprintln(w, "Hello")
	case <- ctx.Done():
		err := ctx.Err()
		log.Println(ctx, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

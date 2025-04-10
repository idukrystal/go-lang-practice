package log

import (
	"context"
	"net/http"
	"log"
	"math/rand"
)

type key int
const requestIdKey key = 42

func Decorator(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := rand.Int63()
		
		ctx := r.Context()
		ctx = context.WithValue(ctx, requestIdKey, id)

		f(w, r.WithContext(ctx))
	}	
}

func Println(ctx context.Context, msg string) {
	id, ok := ctx.Value(requestIdKey).(int64)
	if !ok {
		log.Fatal("id not found")
		return
	}
	log.Printf("%d: %s", id, msg)
}
	

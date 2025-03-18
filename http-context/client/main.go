package main

import (
	"io"
	"log"
	"context"
	"net/http"
	"os"
	"time"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
	ctx = context.WithValue(ctx, "foo", "bar")
	req = req.WithContext(ctx)
	fmt.Printf("value of foo: %v\n", req.Context().Value("foo"))
	
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Println(res.StatusCode)
	}
	io.Copy(os.Stdout, res.Body)
}

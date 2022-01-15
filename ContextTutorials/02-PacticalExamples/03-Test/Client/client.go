package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Println("client")

	rootCtx := context.Background()
	req, err := http.NewRequest("GET", "http://127.0.0.1:8989", nil)
	if err != nil {
		panic(err)
	}

	// create context
	ctx, cancel := context.WithTimeout(rootCtx, 500*time.Second)
	defer cancel()

	// attach context to our request
	req = req.WithContext(ctx)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error ", err.Error())
	}

	log.Println("resp received", resp)
}

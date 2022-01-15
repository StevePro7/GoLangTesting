package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("client")

	req, err := http.NewRequest("GET", "http://127.0.0.1:8989", nil)
	if err != nil {
		panic(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	log.Println("resp received", resp)
}

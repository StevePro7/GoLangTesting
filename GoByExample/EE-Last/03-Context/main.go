package main

import (
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Println("SERVER:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)
}

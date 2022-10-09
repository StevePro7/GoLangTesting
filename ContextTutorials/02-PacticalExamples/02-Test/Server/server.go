package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Println("server")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("request received")
		time.Sleep(time.Second * 3)
		_, err := fmt.Fprintf(w, "Response")
		if err != nil {
			panic(err)
		} // send data to client side
		log.Println("response sent")
	})

	err := http.ListenAndServe("127.0.0.1:8989", nil) // set listen port
	if err != nil {
		panic(err)
	}
}

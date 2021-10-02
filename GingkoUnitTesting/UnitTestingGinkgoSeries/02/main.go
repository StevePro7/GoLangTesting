package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Sum(x, y int) int {
	return x + y
}

func Handler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hi there endpoint %s", r.URL.Path[1:])
	if err != nil {
		return
	}
}

func ReadHandler(w http.ResponseWriter, _ *http.Request) {
	dat, err := ioutil.ReadFile("data.txt")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = fmt.Fprintf(w, "Content in file is \r\n %s", string(dat))
	if err != nil {
		return
	}
}

func main() {
	http.HandleFunc("/", Handler)
	http.HandleFunc("/read", ReadHandler)
	err := http.ListenAndServe(":8002", nil)
	if err != nil {
		return
	}
}

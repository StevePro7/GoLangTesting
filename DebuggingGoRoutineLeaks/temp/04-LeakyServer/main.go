package main

import (
	"fmt"
	"github.com/rpccloud/goid"
	"log"
	"net/http"
	"strconv"
	"time"
)

// function to add an array of numbers
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		id := goid.GetRoutineId()
		fmt.Println(id, v)
		time.Sleep(1 * time.Second)
		sum += v
	}
	// writes the sum to the go routines
	c <- sum // send sum to c
}

func sumConcurrent(w http.ResponseWriter, _ *http.Request) {
	id := goid.GetRoutineId()

	s := []int{7, 2, 8, -9, 4, 0}

	c1 := make(chan int)
	c2 := make(chan int)

	// spin up a goroutine
	go sum(s[:len(s)/2], c1)

	// spin up a goroutine
	go sum(s[len(s)/2:], c2)

	x := <-c1

	var err error
	_, err = fmt.Fprintf(w, strconv.Itoa(int(id)))
	_, err = fmt.Fprintf(w, " => ")
	_, err = fmt.Fprintf(w, strconv.Itoa(x))

	if err != nil {
		_, err := fmt.Fprintln(w, err.Error())
		if err != nil {
			return
		}
	}
}
func main() {

	http.HandleFunc("/sum", sumConcurrent)
	err := http.ListenAndServe(":8002", nil)
	if err != nil {
		log.Fatal("L+S ", err)
	}
}

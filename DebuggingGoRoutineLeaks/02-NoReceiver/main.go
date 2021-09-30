package main

import (
	fmt "fmt"
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

func sumConcurrent(w http.ResponseWriter, r *http.Request) {
	id := goid.GetRoutineId()

	s := []int{7, 2, 8, -9, 4, 0}

	c1 := make(chan int)
	c2 := make(chan int)

	// spin up a goroutine
	go sum(s[:len(s)/2], c1)

	// spin up a goroutine
	go sum(s[len(s)/2:], c2)

	x := <-c1
	fmt.Fprintf(w, strconv.Itoa(int(id)))
	fmt.Fprintf(w, " => ")
	fmt.Fprintf(w, strconv.Itoa(x))

}
func main() {

	http.HandleFunc("/sum", sumConcurrent)
	err := http.ListenAndServe(":8002", nil)
	if err != nil {
		log.Fatal("L+S ", err)
	}
}

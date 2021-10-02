package main

import (
	"fmt"
	"github.com/rpccloud/goid"
	"log"
	"net/http"
	"runtime"
	"runtime/debug"
	"runtime/pprof"
	"strconv"
)

// get the count of number of go routines in the system
func countGoRoutines() int {
	return runtime.NumGoroutine()
}

// respond with number of go routines in the system
func getGoroutineCountHandler(w http.ResponseWriter, _ *http.Request) {
	// Get the count of number of go routines running
	count := countGoRoutines()
	_, err := w.Write([]byte(strconv.Itoa(count)))
	if err != nil {
		return
	}
}

// respond with the stack trace of the system
func getStackTraceHandler(w http.ResponseWriter, _ *http.Request) {
	stack := debug.Stack()
	_, err := w.Write([]byte("stevepro01"))
	if err != nil {
		return
	}
	_, err = w.Write([]byte("stevepro02"))
	_, err = w.Write(stack)
	if err != nil {
		return
	}

	_, err = w.Write([]byte("stevepro03"))
	err = pprof.Lookup("goroutine").WriteTo(w, 2)
	if err != nil {
		log.Fatal(err)
	}

	_, err = w.Write([]byte("stevepro04"))
}

// function to add an array of numbers
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	// writes the sum to the go routines
	c <- sum // send sum to c
}

// HTTP handler for /sum
func sumConcurrent(w http.ResponseWriter, _ *http.Request) {
	s := []int{7, 2, 8, -9, 4, 0}

	c1 := make(chan int)
	c2 := make(chan int)

	// spin up a goroutine
	go sum(s[:len(s)/2], c1)

	// spin up a goroutine
	go sum(s[len(s)/2:], c2)

	// not reading from c2
	// go routine writing to c2 will be blocked
	x := <-c1

	// write the response
	_, err := fmt.Fprintf(w, strconv.Itoa(x))
	if err != nil {
		return
	}
}

func main() {
	id := goid.GetRoutineId()
	fmt.Println(id)

	// get the sum of numbers
	http.HandleFunc("/sum", sumConcurrent)

	// get the count of number of go routines in the system
	http.HandleFunc("/count", getGoroutineCountHandler)

	// respond with the stack trace of the system
	http.HandleFunc("/stack", getStackTraceHandler)

	err := http.ListenAndServe(":8002", nil)
	if err != nil {
		log.Fatal("ListenAndServer : ", err)
	}
}

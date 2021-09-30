package main

import (
	"fmt"
	"github.com/rpccloud/goid"
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

func main() {
	id := goid.GetRoutineId()
	fmt.Println(id, "start")

	s := []int{7, 2, 8, -9, 4, 0}

	c1 := make(chan int)
	c2 := make(chan int)

	// spin up a goroutine
	go sum(s[:len(s)/2], c1)

	// spin up a goroutine
	go sum(s[len(s)/2:], c2)

	x := <-c1
	fmt.Println(id, x)
	y := <-c2
	fmt.Println(id, y)

	fmt.Println(id, x+y)
	fmt.Println(id, "-end-")
}

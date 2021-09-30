package main

import (
	"fmt"
)

// function to add an array of numbers
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		fmt.Println(v)
		sum += v
	}
	// writes the sum to the go routines
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c1 := make(chan int)
	c2 := make(chan int)

	// spin up a goroutine
	go sum(s[:len(s)/2], c1)

	// spin up a goroutine
	go sum(s[len(s)/2:], c2)

	// receive from c1 and c2
	x, y := <-c1, <-c2
	sum := x + y

	fmt.Println(x, y, sum)
}

package main

import (
	"fmt"

	"github.com/nikandfor/goid"
)

func f(left, right chan int) {
	left <- 1 + <-right
}
func main() {
	fmt.Printf("[%d] beg\n", goid.ID())
	const n = 100000
	leftmost := make(chan int)
	right := leftmost
	left := leftmost
	for i := 0; i < n; i++ {
		right = make(chan int)
		go f(left, right)
		left = right
	}
	go func(c chan int) {
		c <- 1
	}(right)
	fmt.Printf("[%d] %d\n", goid.ID(), <-leftmost)
	fmt.Printf("[%d] end\n", goid.ID())
}

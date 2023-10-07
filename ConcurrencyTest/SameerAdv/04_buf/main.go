package main

import (
	"fmt"

	"github.com/nikandfor/goid"
)

type Ball struct {
	hits int
}

func main() {
	fmt.Printf("[%d] beg\n", goid.ID())
	in, out := make(chan int), make(chan int)
	go buffer(in, out)
	for i := 0; i < 10; i++ {
		in <- i
	}
	close(in)
	for i := range out {
		fmt.Printf("[%d] %d\n", goid.ID(), i)
	}
	fmt.Printf("[%d] end\n", goid.ID())
}

// buffer provides an unbounded buffer between in and out.
// buffer exits when it is closed and all items in the
// buffer have been sent to out at which point closes
func buffer(in <-chan int, out chan<- int) {
	var buf []int
	for in != nil || len(buf) > 0 {
		var i int
		var c chan<- int
		if len(buf) > 0 {
			i = buf[0]
			c = out // enable send case
		}
		select {
		case n, ok := <-in:
			if ok {
				buf = append(buf, n)
			} else {
				in = nil // disable reseive case
			}
		case c <- i:
			buf = buf[1:]
		}
	}
	close(out)
}

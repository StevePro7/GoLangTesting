package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/nikandfor/goid"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	fmt.Printf("[%d] beg\n", goid.ID())

	a, b := make(chan string), make(chan string)
	go func() { a <- "a" }()
	go func() { b <- "b" }()
	if rand.Intn(2) == 0 {
		a = nil
		fmt.Printf("[%d] nil a\n", goid.ID())
	} else {
		b = nil
		fmt.Printf("[%d] nil b\n", goid.ID())
	}

	select {
	case s := <-a:
		fmt.Printf("[%d] A got %s\n", goid.ID(), s)
	case s := <-b:
		fmt.Printf("[%d] B got %s\n", goid.ID(), s)
	}

	fmt.Printf("[%d] end\n", goid.ID())
}

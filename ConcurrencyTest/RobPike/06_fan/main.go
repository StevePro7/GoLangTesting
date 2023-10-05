package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/nikandfor/goid"
)

func fanin(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			val1 := <-input1
			c <- val1
		}
	}()
	go func() {
		for {
			val2 := <-input2
			c <- val2
		}
	}()

	return c
}

// returns receive-only channel of strings
func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			val := fmt.Sprintf("[%d] %s %d", goid.ID(), msg, i)
			c <- val
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func main() {
	fmt.Printf("[%d] beg\n", goid.ID())
	c := fanin(boring("joe"), boring("ann"))
	for i := 0; i < 7; i++ {
		val := <-c
		fmt.Printf("[%d] %s\n", goid.ID(), val)
	}
	fmt.Printf("[%d] end\n", goid.ID())
}

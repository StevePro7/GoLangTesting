package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/nikandfor/goid"
)

func boring(msg string, c chan string) {
	for i := 0; ; i++ {
		val := fmt.Sprintf("[%d] %s %d", goid.ID(), msg, i)
		c <- val
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func main() {
	c := make(chan string)
	fmt.Printf("[%d] beg\n", goid.ID())
	go boring("boring", c)
	for i := 0; i < 5; i++ {
		val := <-c
		fmt.Printf("[%d] say: %q\n", goid.ID(), val)
	}
	fmt.Printf("[%d] end\n", goid.ID())
}

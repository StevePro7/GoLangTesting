package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/nikandfor/goid"
)

// returns receive-only channel of strings
func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			val := fmt.Sprintf("[%d] %s %d", goid.ID(), msg, i+1)
			c <- val
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func main() {
	fmt.Printf("[%d] beg\n", goid.ID())
	c := boring("borring")
	for i := 0; i < 5; i++ {
		val := <-c
		fmt.Printf("[%d] say %q\n", goid.ID(), val)
	}
	fmt.Printf("[%d] end\n", goid.ID())
}

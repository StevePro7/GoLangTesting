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
		for i := 0; i < 5; i++ {
			val := fmt.Sprintf("[%d] %s %d", goid.ID(), msg, i+1)
			c <- val
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func main() {

	fmt.Printf("[%d] beg\n", goid.ID())
	done := make(chan bool)
	c := boring("Joe")
	go func() {
		for {
			select {
			case s := <-c:
				fmt.Printf("[%d] %s\n", goid.ID(), s)
			case <-time.After(time.Duration(1 * time.Second)):
				fmt.Printf("[%d] You're too slow!\n", goid.ID())
				done <- true
			}
		}
	}()
	fmt.Printf("[%d] BLK\n", goid.ID())
	<-done
	fmt.Printf("[%d] end\n", goid.ID())
}

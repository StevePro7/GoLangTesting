package main

import (
	"fmt"
	"math/rand"

	"github.com/nikandfor/goid"
)

// returns receive-only channel of strings
func boring(msg string, quit chan bool) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case c <- fmt.Sprintf("[%d] do nothing\n", goid.ID()):
				// dp nothing
			case <-quit:
				fmt.Printf("[%d] quitting out!\n", goid.ID())
				return
			}
		}
	}()
	return c
}

func main() {

	fmt.Printf("[%d] beg\n", goid.ID())
	quit := make(chan bool)
	c := boring("Joe", quit)
	for i := rand.Intn(10); i >= 0; i-- {
		fmt.Println(<-c)
	}
	quit <- true
	fmt.Printf("[%d] end\n", goid.ID())
}

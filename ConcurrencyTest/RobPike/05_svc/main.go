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
			c <- fmt.Sprintf("[%d] %s %d", goid.ID(), msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func main() {
	joe := boring("joe")
	ann := boring("ann")
	for i := 0; i < 5; i++ {
		fmt.Printf("[%d] say %q\n", goid.ID(), <-joe)
		fmt.Printf("[%d] say %q\n", goid.ID(), <-ann)
	}
	fmt.Println("end")
}

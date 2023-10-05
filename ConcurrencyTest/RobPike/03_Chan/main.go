package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/nikandfor/goid"
)

func boring(msg string, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("[%d] %s %d", goid.ID(), msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func main() {
	c := make(chan string)
	fmt.Println(goid.ID(), "beg")
	go boring("boring", c)
	for i := 0; i < 5; i++ {
		fmt.Printf("say: %q\n", <-c)
	}
	fmt.Println(goid.ID(), "end")
}

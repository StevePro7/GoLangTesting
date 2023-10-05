package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/nikandfor/goid"
)

func boring(msg string) {
	for i := 0; ; i++ {
		fmt.Printf("[%d] %s %d\n", goid.ID(), msg, i)
		//time.Sleep(time.Second)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func main() {
	fmt.Printf("[%d] beg\n", goid.ID())
	go boring("boring")
	time.Sleep(2 * time.Second)
	fmt.Printf("[%d] end\n", goid.ID())
}

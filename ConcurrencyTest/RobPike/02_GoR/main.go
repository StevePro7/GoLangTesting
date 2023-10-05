package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/nikandfor/goid"
)

func boring(msg string) {
	for i := 0; ; i++ {
		fmt.Println(goid.ID(), msg, i)
		//time.Sleep(time.Second)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func main() {
	go boring("boring")
	fmt.Println(goid.ID(), "listen")
	time.Sleep(2 * time.Second)
	fmt.Println(goid.ID(), "end")
}

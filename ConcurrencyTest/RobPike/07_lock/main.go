package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/nikandfor/goid"
)

type Message struct {
	str  string
	wait chan bool
}

func fanin(input1, input2 <-chan Message) <-chan Message {
	c := make(chan Message)
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
func boring(msg string) <-chan Message {
	waitForIt := make(chan bool) // Shared between all messages
	c := make(chan Message)
	go func() {
		for i := 0; ; i++ {
			str := fmt.Sprintf("[%d] %s %d", goid.ID(), msg, i)
			val := Message{str, waitForIt}
			c <- val
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			<-waitForIt
		}
	}()
	return c
}

func main() {

	fmt.Printf("[%d] beg\n", goid.ID())
	//c := make(chan Message)
	c := fanin(boring("joe"), boring("ann"))
	for i := 0; i < 5; i++ {
		msg1 := <-c
		fmt.Printf("[%d] %s\n", goid.ID(), msg1.str)
		msg2 := <-c
		fmt.Printf("[%d] %s\n", goid.ID(), msg2.str)
		val := true
		msg1.wait <- val
		msg2.wait <- val

	}
	fmt.Printf("[%d] end\n", goid.ID())
}

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Message struct {
	str  string
	wait chan bool
}

func boring(c chan Message) {
	for i := 0; i < 5; i++ {
	}
}
func main() {

	waitForIn := make(chan bool)
	c := make(chan Message)
	go boring(c)
	msg := "test"
	i := 0
	c <- Message{
		fmt.Sprintf("%s %d", msg, i),
		waitForIn,
	}
	time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	<-waitForIn
	fmt.Println("end")
}

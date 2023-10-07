package main

import (
	"fmt"
	"time"

	"github.com/nikandfor/goid"
)

type Ball struct {
	hits int
}

func main() {
	fmt.Printf("[%d] beg\n", goid.ID())
	table := make(chan *Ball)
	go player("ping", table)
	go player("pong", table)

	table <- new(Ball)
	time.Sleep(1 * time.Second)
	<-table

	fmt.Printf("[%d] end\n", goid.ID())
	panic("show the stacks")
}

func player(name string, table chan *Ball) {
	for {
		recd := <-table
		ball := recd
		ball.hits++
		fmt.Printf("[%d] %s %d\n", goid.ID(), name, ball.hits)
		time.Sleep(100 * time.Millisecond)
		send := ball
		table <- send
	}
}

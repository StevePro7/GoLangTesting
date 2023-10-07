package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/nikandfor/goid"
)

type Result struct {
	text string
}
type Search func(query string) Result

var (
	Web   = fakeSearch("web")
	Image = fakeSearch("image")
	Video = fakeSearch("video")
)

func First(query string, replicas ...Search) Result {
	c := make(chan Result)
	searchReplica := func(i int) {
		send := replicas[i](query)
		c <- send
	}
	for i := range replicas {
		go searchReplica(i)
	}
	recd := <-c
	return recd
}

func fakeSearch(kind string) Search {
	return func(query string) Result {
		duration := time.Duration(rand.Intn(100)) * time.Millisecond
		time.Sleep(duration)
		msg := fmt.Sprintf("[%d] %s result for %q\n", goid.ID(), kind, query)
		return Result{msg}
	}
}

func main() {
	fmt.Printf("[%d] beg\n", goid.ID())
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	result := First("golang", fakeSearch("replica1"), fakeSearch("replica2"))
	elapsed := time.Since(start)
	fmt.Printf("[%d] %q\n", goid.ID(), result)
	fmt.Printf("[%d] %q\n", goid.ID(), elapsed)
	fmt.Printf("[%d] end\n", goid.ID())
}

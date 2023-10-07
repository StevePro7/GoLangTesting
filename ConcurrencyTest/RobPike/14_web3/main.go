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

func Google(query string) (results []Result) {
	c := make(chan Result)
	go func() { c <- Web(query) }()
	go func() { c <- Image(query) }()
	go func() { c <- Video(query) }()

	timeout := time.After(80 * time.Millisecond)
	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			fmt.Printf("[%d] timed out\n", goid.ID())
			return
		}
	}
	return
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
	results := Google("golang")
	elapsed := time.Since(start)
	fmt.Printf("[%d] %q\n", goid.ID(), results)
	fmt.Printf("[%d] %q\n", goid.ID(), elapsed)
	fmt.Printf("[%d] end\n", goid.ID())
}

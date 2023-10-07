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

	fmt.Printf("[%d] end\n", goid.ID())
}

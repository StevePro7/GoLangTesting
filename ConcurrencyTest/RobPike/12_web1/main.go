package main

import (
	"crypto/rand"
	"fmt"
	"time"
)

type Result struct{}
type Search func(query string) Result

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100) * time.Millisecond)
		return Result
	}
}
func main() {
	fmt.Println("bob")
}

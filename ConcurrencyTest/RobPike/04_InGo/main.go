package main

import (
	"fmt"
	"math/rand"
	"time"
)

// returns receive-only channel of strings
func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			val := fmt.Sprintf("%s %d", msg, i+1)
			c <- val
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func main() {
	c := boring("borring")
	for i := 0; i < 5; i++ {
		val := <-c
		fmt.Printf("say %q\n", val)
	}
	fmt.Println("end")
}

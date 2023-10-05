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
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func main() {
	joe := boring("joe")
	ann := boring("ann")
	for i := 0; i < 5; i++ {
		fmt.Printf("say %q\n", <-joe)
		fmt.Printf("say %q\n", <-ann)
	}
	fmt.Println("end")
}

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// https://gobyexample.com/signals
func main() {

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	done := make(chan bool, 1)

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Printf("sgb goroutine:'%s'\n", sig)
		// send true to the done channel
		//done <- false
		//done <- true
		close(done) // false
	}()

	fmt.Println("awaiting signal")
	// block!!
	// receive bool from the done channel
	ch := <-done
	fmt.Printf("exiting [%v]", ch)
}

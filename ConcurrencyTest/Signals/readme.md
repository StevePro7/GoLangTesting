Signals
27-Sep-2023

https://gobyexample.com/signals

VS Code
go run main.go
Ctrl+C
interrupt

Terminal
go run main.go
Click X to exit
terminated

However if
done <- false
OR
close(done)

then Click X to exit won't print terminated

IMPORTANT
right at the end the code must be
fmt.Println("exiting")

otherwise terminated will not be printed
on click X to exit 





old code


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
		done <- true
		//close(done) // false
	}()

	fmt.Println("awaiting signal")
	// block!!
	// receive bool from the done channel
	//ch := <-done
	//ch := 0
	<-done
	fmt.Printf("exiting")
	//fmt.Printf("exiting [%v]", ch)
}

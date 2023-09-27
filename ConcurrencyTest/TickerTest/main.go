package main

import (
	"fmt"
	"time"
)

func main() {
	Ticker := time.NewTicker(2 * time.Second)
	mychannel := make(chan bool)

	go func() {

		for {
			select {
			case <-mychannel:
				return
			case tm := <-Ticker.C:
				fmt.Println("The Current time is: ", tm)
			}
		}
	}()

	time.Sleep(7 * time.Second)
	Ticker.Stop()
	mychannel <- true
	fmt.Println("the end")
	fmt.Println("the end2")
}

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func doWork(ctx context.Context, resChan chan int) {
	log.Println("[doWork] launch the doWork")
	sum := 0
	for {
		log.Println("[doWork] one iteration")
		time.Sleep(time.Millisecond)
		select {
		case <-ctx.Done():
			log.Println("[doWork] ctx Done is received inside doWork")
			return
		default:
			sum++
			time.Sleep(10 * time.Minute)
			if sum > 1000 {
				log.Println("[doWork] sum has reached 1000")
				resChan <- sum
				return
			}
		}
	}
}

func main() {
	fmt.Println("server")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("[Handler] request received")

		// retrieve the context of th request
		rCtx := r.Context()

		// create the result channel
		resChan := make(chan int)

		// launch the function doWork on a goroutine
		go doWork(rCtx, resChan)

		// Wait for
		// 1. the client drops the connection
		// 2. the function doWork to finish it works
		select {
		case <-rCtx.Done():
			log.Println("[Handler] context canceled in main handler, client disconnected")
			return
		case result := <-resChan:
			log.Println("[Handler] received 1000")
			log.Println("[Handler] send response")
			_, err := fmt.Fprintf(w, "Response %d", result)
			if err != nil {
				panic(err)
			} // send data to client side
			return
		}
	})

	err := http.ListenAndServe("127.0.0.1:8989", nil) // set listen port
	if err != nil {
		panic(err)
	}
}

package main

import (
	"context"
	"log"
	"runtime"
	"time"
)

func doSth1(ctx context.Context) {
	select {
	case <-ctx.Done():
		log.Println("1st goroutine return")
		return
	}
}

func doSth2(ctx context.Context) {
	select {
	case <-ctx.Done():
		log.Println("2nd goroutine return")
		return
	}
}

func launch() {
	ctx := context.Background()

	// Goroutine count: 3
	//ctx, _ = context.WithCancel(ctx)

	// Goroutine count: 1
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	log.Println("launch 1st goroutine")
	go doSth1(ctx)
	log.Println("launch 2nd goroutine")
	go doSth2(ctx)
}

func main() {
	log.Println("begin program")
	go launch()
	time.Sleep(time.Millisecond)
	log.Printf("Goroutine count: %d\n", runtime.NumGoroutine())
	for {
	}
}

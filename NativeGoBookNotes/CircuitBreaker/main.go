package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Circuit func(context.Context) (string, error)

func Breaker(circuit Circuit, fialureThreshold uint) Circuit {
	var consecutiveFailures int = 0
	var lastAttempt  time.Now()
	var m sync.RWMutex


}


func main() {

	fmt.Println("bc")
}

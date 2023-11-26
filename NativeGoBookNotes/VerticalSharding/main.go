package main

import (
	"fmt"
	"sync"
)

type unshardedMap struct {
	sync.RWMutex
	m map[string]interface{}
}

func newUnshardedMap() *unshardedMap {
	return &unshardedMap{
		m: make(map[string]interface{}),
	}
}

func main() {
	fmt.Println("hello??")
}

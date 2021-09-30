package main

import (
	"fmt"
	"github.com/rpccloud/goid"
	"log"
	"net/http"
	"runtime/debug"
	"runtime/pprof"
)

func getStackTraceHandler(w http.ResponseWriter, _ *http.Request) {
	stack := debug.Stack()
	_, err := w.Write([]byte("stevepro01"))
	if err != nil {
		return
	}
	_, err = w.Write([]byte("stevepro02"))
	_, err = w.Write(stack)
	if err != nil {
		return
	}

	_, err = w.Write([]byte("stevepro03"))
	err = pprof.Lookup("goroutine").WriteTo(w, 2)
	if err != nil {
		log.Fatal(err)
	}

	_, err = w.Write([]byte("stevepro04"))
}

func main() {
	id := goid.GetRoutineId()
	fmt.Println(id)
	http.HandleFunc("/stack", getStackTraceHandler)

	err := http.ListenAndServe(":8002", nil)
	if err != nil {
		log.Fatal(err)
	}
}

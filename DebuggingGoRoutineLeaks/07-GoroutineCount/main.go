package main

import (
	"fmt"
	"github.com/rpccloud/goid"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"
)

const (
	leakyServer = "http://localhost:8002"
)

// Send get request and return the response body
func getReq(endPoint string) ([]byte, error) {
	response, err := http.Get(leakyServer + endPoint)
	if err != nil {
		return []byte{}, err
	}

	defer func(body io.ReadCloser) {
		err := body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(response.Body)

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return []byte{}, err
	}
	return body, nil
}

// get the count of the number of go routines in the server
func getRoutineCount() (int, error) {
	body, err := getReq("/count")
	if err != nil {
		return -1, err
	}

	count, err := strconv.Atoi(string(body))
	if err != nil {
		return -1, err
	}

	return count, nil
}

// obtain stack trace of the server
func getStackTrace() (string, error) {
	body, err := getReq("/stack")
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func main() {
	id := goid.GetRoutineId()
	fmt.Println(id)

	// get the number of go routines in the leaky server
	count, err := getRoutineCount()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("\n %d go routines before the load test in the system.", count)

	var wg sync.WaitGroup
	// send 50 concurrent request to the leaky endpoint
	for i := 0; i < 50; i++ {
		wg.Add(i)
		go func() {
			defer wg.Done()

			id := goid.GetRoutineId()
			fmt.Println(id)

			_, err = getReq("/sum")
			if err != nil {
				log.Fatal(err)
			}
		}()
	}
	wg.Wait()

	// get the count of number of goroutines in the system after the load test
	count, err = getRoutineCount()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("\n %d go routines after the load tset inthe system.", count)

	// obtain the stack trace of the system
	trace, err := getStackTrace()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("\n stack trace after the load test : \n %s", trace)
}

package main

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hi there, the end point is :%s!", r.URL.Path[1:])
	if err != nil {
		return
	}
}

func ReadHandler(w http.ResponseWriter, _ *http.Request) {
	dat, err := ioutil.ReadFile("data.txt")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = fmt.Fprintf(w, "Content in file is...\r\n%s", string(dat))
	if err != nil {
		return
	}
}

func main() {
	url := "http://localhost:8002/read"
	body, err := getResponse(url)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

func getResponse(url string) ([]byte, error) {
	if len(url) == 0 {
		return nil, errors.New("invalid url")
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	c := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	code := resp.StatusCode
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil && code != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}

	if code != http.StatusOK {
		return nil, fmt.Errorf("server status error: %v",
			http.StatusText(code))
	}

	return body, nil
}

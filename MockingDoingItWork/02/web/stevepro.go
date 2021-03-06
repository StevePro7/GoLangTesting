package web

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var (
	Client HTTPClient
)

func init() {
	Client = &http.Client{}
}

func GetFixedValue(baseURL string) (string, error) {
	url := fmt.Sprintf("%s", baseURL)

	request, _ := http.NewRequest(http.MethodGet, url, nil)
	request.Header.Add("Accept", "application/json")
	client := &http.Client{}

	response, err := client.Do(request)

	if response.StatusCode != 200 {
		return "", err
	}

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	v := &ValueHolder{}

	err = json.Unmarshal(content, v)
	if err != nil {
		return "", err
	}

	return v.Value, nil
}

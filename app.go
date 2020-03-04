package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	counter := 0
	// Configure number of reqs
	for i := 0; i < 1000000000; i++ {
		counter += i

		response, err := request()
		s := string(response)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(s)
	}
}

func request() ([]byte, error) {
	resp, err := http.Post("http://localhost:5000", "text/plain", nil)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	return body, nil
}

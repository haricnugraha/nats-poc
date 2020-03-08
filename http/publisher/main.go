package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func main() {
	subsriberURL := "http://localhost:8080"
	contentType := "application/json"

	startTime := time.Now()

	// send request
	resp, err := http.Post(subsriberURL, contentType, strings.NewReader(`{id: "random-id", total: 2000000}`))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer resp.Body.Close()
	duration := time.Since(startTime).Microseconds()
	fmt.Printf("Duration: %d microseconds \n", duration)

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(responseBody))
}

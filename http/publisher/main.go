package main

import (
	"io/ioutil"
	"log"
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
		log.Println(err.Error())
		return
	}
	defer resp.Body.Close()
	duration := time.Since(startTime).Microseconds()
	log.Printf("Duration: %d microseconds \n", duration)

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println(string(responseBody))
}

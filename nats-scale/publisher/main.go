package main

import (
	"log"
	"strconv"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	// connect to nats server
	nc, err := nats.Connect(nats.DefaultURL)
	// Close connection
	defer nc.Close()
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println("Connected to " + nats.DefaultURL)

	subject := "order.paid"

	numOfMessages := 10
	for i := 0; i < numOfMessages; i++ {
		id := strconv.Itoa(i)
		message := []byte(`{id: ` + id + `, total: 2000000}`)
		startTime := time.Now()

		// publish message
		nc.Publish(subject, message)
		duration := time.Since(startTime).Microseconds()
		log.Printf("ID: %s, Duration: %d microseconds \n", id, duration)
	}
}

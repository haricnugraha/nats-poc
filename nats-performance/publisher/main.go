package main

import (
	"log"
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
	message := []byte(`{id: random-id, total: 2000000}`)

	// publish message
	startTime := time.Now()
	nc.Publish(subject, message)
	duration := time.Since(startTime).Microseconds()
	log.Printf("Duration: %d microseconds \n", duration)
}

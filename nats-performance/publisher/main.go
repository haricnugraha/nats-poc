package main

import (
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	// connect to nats server
	nc, err := nats.Connect(nats.DefaultURL)
	fmt.Println("Connected to " + nats.DefaultURL)
	// Close connection
	defer nc.Close()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	subject := "order.paid"
	message := []byte(`{id: random-id, total: 2000000}`)
	startTime := time.Now()

	// publish message
	nc.Publish(subject, message)
	duration := time.Since(startTime).Microseconds()
	fmt.Printf("Duration: %d microseconds \n", duration)
}

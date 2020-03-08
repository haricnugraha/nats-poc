package main

import (
	"fmt"

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

	// Simple Async Subscriber
	nc.QueueSubscribe("order.*", "orderWorker", func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})

	// just prevent program exited
	wait := make(chan struct{})
	<-wait
}

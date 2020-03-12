package main

import (
	"log"

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

	// Simple Async Subscriber
	nc.QueueSubscribe("order.*", "orderWorker", func(m *nats.Msg) {
		log.Printf("Received a message: %s\n", string(m.Data))
	})

	// just prevent program exited
	wait := make(chan struct{})
	<-wait
}

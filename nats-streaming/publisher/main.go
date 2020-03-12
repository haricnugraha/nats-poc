package main

import (
	"log"
	"strconv"
	"time"

	"github.com/nats-io/stan.go"
)

const (
	clusterID = "test-cluster"
	clientID  = "order-service-1"
	natsURL   = "nats://localhost:5555"
	subject   = "order.paid"
)

func main() {
	// Connect to NATS Streaming server
	sc, err := stan.Connect(
		clusterID,
		clientID,
		stan.NatsURL(natsURL),
		// stan.MaxPubAcksInflight(10),
	)
	defer sc.Close()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Connected to nats streaming server")

	numOfMessages := 10
	for i := 0; i < numOfMessages; i++ {
		id := strconv.Itoa(i)
		message := []byte(`{id: ` + id + `, total: 2000000}`)
		startTime := time.Now()

		ackHandler := func(ackedNuid string, err error) {
			if err != nil {
				log.Printf("Error publishing message id %s: %v\n", ackedNuid, err.Error())
				return
			}

			log.Printf("Received ACK for message id %s\n", ackedNuid)
		}

		// publish message
		messageGUID, err := sc.PublishAsync(subject, message, ackHandler)
		if err != nil {
			log.Printf("Error: %s", err.Error())
		}
		duration := time.Since(startTime).Microseconds()
		log.Printf("Order ID: %s, Message ID: %s, Duration: %d microseconds \n", id, messageGUID, duration)
	}

	// just prevent program exited
	wait := make(chan struct{})
	<-wait
}

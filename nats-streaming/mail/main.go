package main

import (
	"log"

	"github.com/nats-io/stan.go"
)

const (
	clusterID = "test-cluster"
	clientID  = "mail-service-1"
	natsURL   = "nats://localhost:5555"
	subject   = "order.paid"
	durableID = "mail-service-durable"
)

func main() {
	// Connect to NATS Streaming server
	sc, err := stan.Connect(
		clusterID,
		clientID,
		stan.NatsURL(natsURL),
	)
	defer sc.Close()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Connected to nats streaming server")

	// Subscribe with manual ack mode, and set AckWait to 60 seconds
	// aw, _ := time.ParseDuration("60s")
	sc.Subscribe(subject, func(msg *stan.Msg) {
		// Manual ACK
		// msg.Ack()
		// Handle the message
		log.Printf("Subscribed message from clientID - %s for Order: %s\n", clientID, string(msg.Data))
	},
		// set durable id
		stan.DurableName(durableID),

	// starting position
	// stan.StartWithLastReceived())
	// stan.DeliverAllAvailable(),
	// stan.StartAtSequence(22)
	// stan.StartAtTime(time.Now())

	// stan.MaxInflight(25),

	// stan.SetManualAckMode(),
	// stan.AckWait(aw),
	)

	// qsub2, _ := sc.QueueSubscribe(channelName,
	// 	queueName, func(m *stan.Msg) {...})

	// unsubscribe, remove durable
	// err := sub.Unsubscribe()

	// just prevent program exited
	wait := make(chan struct{})
	<-wait
}

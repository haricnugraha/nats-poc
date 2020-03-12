# http distributed systems
http-sub:
	go run ./http/subscriber/main.go
http-pub:
	go run ./http/publisher/main.go

# nats distributed systems performance comparison
nats-perf-sub:
	go run ./nats-performance/subscriber/main.go
nats-perf-pub:
	go run ./nats-performance/publisher/main.go

# nats distributed systems scale
nats-scale-inventory:
	go run ./nats-scale/inventory/main.go
nats-scale-inventory2:
	go run ./nats-scale/inventory2/main.go
nats-scale-mail:
	go run ./nats-scale/mail/main.go
nats-scale-pub:
	go run ./nats-scale/publisher/main.go

# nats streaming
nats-stream-mail:
	go run ./nats-streaming/mail/main.go
nats-stream-pub:
	go run ./nats-streaming/publisher/main.go

fmt:
	go vet ./http/** && go fmt ./http/**
	go vet ./nats-performance/** && go fmt ./nats-performance/**
	go vet ./nats-scale/** && go fmt ./nats-scale/**
	go vet ./nats-streaming/** && go fmt ./nats-streaming/**
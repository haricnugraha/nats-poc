# Nats Proof of Concept

Just exploring Nats with Go

## Requirements

- Go. ^1.14
- Docker ^19.03.8

## Usage

### http distributed systems

- Run subscriber

  `make http-sub`

- Run publisher

  `make http-pub`

### Nats performance comparison

- Run subscriber

  `make nats-perf-sub`

- Run publisher

  `make nats-perf-pub`

### Nats scale service

- Run 1st subscriber

  `make nats-scale-inventory`

- Scale 1st subscriber worker

  `make nats-scale-inventory2`

- Add new subscriber service

  `make nats-scale-mail`

- Run publisher

  `make nats-scale-pub`

### Nats streaming

- Run subscriber

  `make nats-stream-mail`

- Run publisher

  `make nats-stream-pub`

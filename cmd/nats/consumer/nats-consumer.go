package main

import (
	"github.com/nats-io/go-nats"
	"github.com/stupendous-man/goctopus/connectors"
	"log"
)

func main() {
	conn := connectors.NatsConnector{NatsUrl: nats.DefaultURL, NatsSubject: "foo.bar"}
	msg := conn.Receive()
	log.Println(msg)
}

package main

import (
	"github.com/nats-io/go-nats"
	"github.com/stupendous-man/goctopus/connectors"
)

func main() {
	conn := connectors.NatsConnector{NatsUrl: nats.DefaultURL, NatsSubject: "foo.bar"}
	conn.Send("We come in peace!")
}

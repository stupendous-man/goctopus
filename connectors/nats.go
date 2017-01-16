package connectors

import (
	"github.com/nats-io/go-nats"
	"log"
	"time"
)

type NatsConnector struct {
	NatsUrl     string
	NatsSubject string
}

func (nc *NatsConnector) Send(request RequestMessage) {
	nconn, err := nats.Connect(nc.NatsUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer nconn.Close()

	nconn.Publish(nc.NatsSubject, []byte(request))
	nconn.Flush()

	if err := nconn.LastError(); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Sent successful")
	}
}

func (nc *NatsConnector) Receive() ResponseMessage {
	nconn, err := nats.Connect(nc.NatsUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer nconn.Close()

	sub, _ := nconn.SubscribeSync(nc.NatsSubject)
	m, err := sub.NextMsg(1 * time.Hour)
	if err != nil {
		panic(err)
	}

	return ResponseMessage(string(m.Data))
}

package connectors

import (
	"github.com/nats-io/go-nats"
	"log"
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

// TODO
/*
func (nc *NatsConnector) Receive() ResponseMessage {

}
*/

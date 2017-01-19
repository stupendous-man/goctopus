package connectors

import "github.com/nsqio/go-nsq"

type NsqConnector struct {
	NsqUrl   string
	NsqTopic string
}

func (nc *NsqConnector) Send(request RequestMessage) {

	producer, _ := nsq.NewProducer(nc.NsqUrl, nsq.NewConfig())

	err := producer.Publish(nc.NsqTopic, []byte(request))
	if err != nil {
		panic(err)
	}

	producer.Stop()
}

func (nc *NsqConnector) Receive() ResponseMessage {

	return nil
}

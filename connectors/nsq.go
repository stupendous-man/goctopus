package connectors

import (
	"github.com/nsqio/go-nsq"
	"sync"
)

type NsqConnector struct {
	NsqUrl     string
	NsqTopic   string
	NsqChannel string
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

	wg := &sync.WaitGroup{}
	wg.Add(1)

	consumer, _ := nsq.NewConsumer(nc.NsqTopic, nc.NsqChannel, nsq.NewConfig())

	var response ResponseMessage

	consumer.AddHandler(
		nsq.HandlerFunc(
			func(message *nsq.Message) error {
				response = ResponseMessage(string(message.Body))
				wg.Done()
				return nil
			}))

	err := consumer.ConnectToNSQLookupd(nc.NsqUrl)
	if err != nil {
		panic(err)
	}

	wg.Wait()

	return response
}

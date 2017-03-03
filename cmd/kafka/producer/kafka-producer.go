package main

import (
	"github.com/stupendous-man/goctopus/connectors"
	"os"
)

func main() {

	conn := connectors.KafkaConnector{KafkaURL: os.Getenv("KAFKA_PORT_9092_TCP_ADDR") + ":" + os.Getenv("KAFKA_PORT_9092_TCP_PORT"), KafkaTopic: "topic"}
	conn.Send("Test Message")
}

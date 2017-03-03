package main

import (
	"github.com/stupendous-man/goctopus/connectors"
	"log"
	"os"
)

func main() {
	conn := connectors.KafkaConnector{KafkaURL: os.Getenv("KAFKA_PORT_9092_TCP_ADDR") + ":" + os.Getenv("KAFKA_PORT_9092_TCP_PORT"), KafkaTopic: "topic"}
	msg := conn.Receive()
	log.Println(msg)
}

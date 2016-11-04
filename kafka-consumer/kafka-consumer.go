package main

import (
	"github.com/Shopify/sarama"
	"log"
	"os"
	"os/signal"
	"time"
)

//Created by stupendous-man
func main() {

	log.Println(os.Getenv("KAFKA_PORT_9092_TCP_ADDR") + ":" + os.Getenv("KAFKA_PORT_9092_TCP_PORT"))
	consumer, err := sarama.NewConsumer([]string{os.Getenv("KAFKA_PORT_9092_TCP_ADDR") + ":" + os.Getenv("KAFKA_PORT_9092_TCP_PORT")}, nil)

	if err != nil {
		//Retry a few times before going into panic
		for retries := 1; retries <= 5; retries++ {
			log.Printf("Retrying kafka connection. Attempt %d...", retries)
			consumer, err = sarama.NewConsumer([]string{os.Getenv("KAFKA_PORT_9092_TCP_ADDR") + ":" + os.Getenv("KAFKA_PORT_9092_TCP_PORT")}, nil)

			if err != nil {
				time.Sleep(100 * time.Millisecond)
			} else {
				break
			}
		}
	}

	if err != nil {
		panic(err)
	}

	//Deferred function to close consumer
	defer func() {
		log.Println("Closing consumer...")
		if err := consumer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	partitionConsumer, err := consumer.ConsumePartition("topic", 0, sarama.OffsetNewest)
	if err != nil {
		panic(err)
	}

	log.Println("Connected to kafka broker...")

	//Deferred function to close partition consumer
	defer func() {
		log.Println("Closing partition consumer...")
		if err := partitionConsumer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	//Trap SIGINT to trigger shutdown
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	//Consume messages from kafka broker on separate goroutine
	go func() {
		for msg := range partitionConsumer.Messages() {
			log.Printf("Consumed message offset :: %d | Value :: %s\n", msg.Offset, string(msg.Value))
		}
	}()

	//Block main thread until SIGINT notification received
	for range signals {
		log.Println("Interrupt signal received. Stopping services...")
		break
	}

}

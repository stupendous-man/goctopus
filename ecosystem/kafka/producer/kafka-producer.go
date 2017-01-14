package main

//TODO: Remove sarama dependency
import (
	"github.com/Shopify/sarama"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"time"
)

//Created by stupendous-man
func main() {

	config := sarama.NewConfig()

	producer, err := sarama.NewAsyncProducer([]string{os.Getenv("KAFKA_PORT_9092_TCP_ADDR") + ":" + os.Getenv("KAFKA_PORT_9092_TCP_PORT")}, config)

	if err != nil {
		//Retry a few times before going into panic
		for retries := 1; retries <= 5; retries++ {
			log.Printf("Retrying kafka connection. Attempt %d...", retries)
			producer, err = sarama.NewAsyncProducer([]string{os.Getenv("KAFKA_PORT_9092_TCP_ADDR") + ":" + os.Getenv("KAFKA_PORT_9092_TCP_PORT")}, config)

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

	//Deferred function to close producer
	defer func() {
		log.Println("Closing producer...")
		producer.AsyncClose()
	}()

	log.Println("Connected to kafka broker...")

	//Trap SIGINT to trigger a graceful shutdown
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	go func() {
		for {
			message := &sarama.ProducerMessage{Topic: "topic", Value: sarama.StringEncoder(generateMessage())}
			producer.Input() <- message
			log.Printf("Produced message :: %s\n", message.Value)
			time.Sleep(500 * time.Millisecond)
		}
	}()

	//Block main thread until SIGINT notification received
	for range signals {
		log.Println("Interrupt signal received. Stopping services...")
		break
	}

}

func random() int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(4)
}

//TODO :: Externalize or expose API to receive messages as input...
func generateMessage() string {
	messages := []string{
		"{Not all who wander are lost...}",
		"{The force is strong with this one...}",
		"{Fortune and glory, kid. Fortune and glory...}",
		"{Why so serious?...}",
	}
	return messages[random()]
}

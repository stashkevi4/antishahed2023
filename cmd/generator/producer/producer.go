package producer

import (
	"github.com/Shopify/sarama"
	"log"
	"strings"
	"time"
)

func ProduceMessage(brokers string, pointId string, text string) {
	// Create a new Kafka configuration
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	// Create a new Kafka producer
	producer, err := sarama.NewSyncProducer(strings.Split(brokers, ","), config)
	if err != nil {
		log.Fatal("Failed to create producer: ", err)
	}
	defer producer.Close()

	// Create a new Kafka message
	message := &sarama.ProducerMessage{
		Topic: "antishahed.engine.sounds",
		Key:   sarama.StringEncoder(pointId + time.UTC.String()),
		Value: sarama.StringEncoder(text), // Replace with your message payload
	}

	// Send the message to Kafka
	partition, offset, err := producer.SendMessage(message)
	if err != nil {
		log.Fatal("Failed to send message: ", err)
	}

	// Print the partition and offset information
	log.Printf("Message sent successfully. Partition: %d, Offset: %d", partition, offset)

}

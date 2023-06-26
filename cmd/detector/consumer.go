package main

import (
	"github.com/Shopify/sarama"
	"log"
	"os"
	"os/signal"
	"strings"
)

func ReadMessage(brokers string) {
	// Create a new Kafka configuration
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	// Create a new Kafka consumer
	consumer, err := sarama.NewConsumer(strings.Split(brokers, "."), config)
	if err != nil {
		log.Fatal("Failed to create consumer: ", err)
	}
	defer consumer.Close()

	// Specify the topic you want to consume from
	topic := "antishahed.engine.sounds" // Replace with your topic name

	// Create a new Kafka partition consumer for the specified topic and partition (or use sarama.OffsetNewest for the latest offset)
	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		log.Fatal("Failed to create partition consumer: ", err)
	}
	defer partitionConsumer.Close()

	// Create a signal channel to handle OS interrupts
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	// Start consuming messages from Kafka
ConsumerLoop:
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			MessageCh <- msg.Value
			log.Printf("Received message: Topic=%s, Partition=%d, Offset=%d, Key=%s, Value=%s\n",
				msg.Topic, msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))

		case err := <-partitionConsumer.Errors():
			log.Println("Error: ", err.Err)

		case <-signals:
			break ConsumerLoop
		}
	}

	log.Println("Consumer stopped")
}

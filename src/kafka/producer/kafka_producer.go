package kafka

import "fmt"

type KafkaProducer struct {
	// Implement Kafka producer logic
	// Use a Kafka library to interact with Kafka broker and publish messages

	string
}

func NewKafkaProducer() *KafkaProducer {
	// Initialize Kafka producer configuration
	// Return a new KafkaProducer instance
	fmt.Println("kafka connected")
	return &KafkaProducer{
		string: "hello",
	}
	//return "hello"
}

func (k *KafkaProducer) Publish(message string) {
	// Publish the message to Kafka topic

	fmt.Println("msg published")
}

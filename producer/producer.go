package main

import (
	"fmt"

	"github.com/Shopify/sarama"
)

func main() {

	server := []string{"localhost:29092"}

	producer, err := sarama.NewSyncProducer(server, nil)

	if err != nil {
		panic(err)
	}

	defer producer.Close()

	msg := &sarama.ProducerMessage{
		Topic:     "test",
		Partition: 0,
		Value:     sarama.StringEncoder("test message"),
	}

	partition, offset, err := producer.SendMessage(msg)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Partition: %d, Offset: %d\n", partition, offset)

}

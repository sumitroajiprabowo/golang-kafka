package main

import "github.com/Shopify/sarama"

func main() {
	server := []string{"localhost:29092"}
	consumer, err := sarama.NewConsumer(server, nil)

	if err != nil {
		panic(err)
	}

	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition("test", 0, sarama.OffsetNewest)

	if err != nil {
		panic(err)
	}

	defer partitionConsumer.Close()

	for {
		select {
		case msg := <-partitionConsumer.Messages():
			println(string(msg.Value))
		}
	}

}

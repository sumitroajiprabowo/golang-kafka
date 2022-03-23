package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/Shopify/sarama"
	"github.com/spf13/viper"
	"github.com/sumitroajiprabowo/golang-kafka/consumer/config"
	"github.com/sumitroajiprabowo/golang-kafka/consumer/events"
	"github.com/sumitroajiprabowo/golang-kafka/consumer/repositories"
	"github.com/sumitroajiprabowo/golang-kafka/consumer/services"
)

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {

	consumer, err := sarama.NewConsumerGroup(viper.GetStringSlice("kafka.server"), viper.GetString("kafka.group"), nil)
	if err != nil {
		panic(err)
	}

	defer consumer.Close()

	db := config.SetupDatabase()
	accountRepo := repositories.NewAccountRepository(db)
	accountEventHandler := services.NewAccountEventHandler(accountRepo)
	accountConsumerHandler := services.NewConsumerHandler(accountEventHandler)

	fmt.Println("Account consumer started...")
	for {
		consumer.Consume(context.Background(), events.Topics, accountConsumerHandler)
	}

}

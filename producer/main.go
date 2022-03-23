package main

import (
	"strings"

	"github.com/Shopify/sarama"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"github.com/sumitroajiprabowo/golang-kafka/producer/controllers"
	"github.com/sumitroajiprabowo/golang-kafka/producer/services"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

func main() {
	producer, err := sarama.NewSyncProducer(viper.GetStringSlice("kafka.servers"), nil)
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	eventProducer := services.NewEventProducer(producer)
	accountService := services.NewAccountService(eventProducer)
	accountController := controllers.NewAccountController(accountService)

	app := fiber.New()

	app.Post("/openaccount", accountController.OpenAccount)
	app.Post("/depositfund", accountController.DepositFund)
	app.Post("/withdrawfund", accountController.WithdrawFund)
	app.Post("/closeaccount", accountController.CloseAccount)

	app.Listen(":8000")
}

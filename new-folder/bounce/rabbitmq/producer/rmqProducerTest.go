package main

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
)

type RabbitMQData struct {
	AccountID  string `json:"account_id"`
	From       string `json:"from"`
	Rcpt       string `json:"rcpt"`
	Domain     string `json:"domain"`
	Event      string `json:"event"`
	BounceType string `json:"bounce_type"`
	Diagnostic string `json:"diagnostic"`
	Retry      int    `json:"retry"`
	IP         string `json:"ip"`
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://rabbit:Maro123!@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"relay_jmta", // name
		"topic",      // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	data := RabbitMQData{AccountID: "1", From: "test@test.com", Rcpt: "test@example.com", Domain: "example.com", Event: "bounced", BounceType: "soft", Diagnostic: "550 not allowed", Retry: 1, IP: "127.0.0.1"}
	body, _ := json.Marshal(data)
	err = ch.Publish(
		"relay_jmta", // exchange
		"#",          // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			Body: body,
		})

	failOnError(err, "Failed to publish a message")

	log.Printf(" [x] Sent %s", body)
}

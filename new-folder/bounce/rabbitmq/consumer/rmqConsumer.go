package consumer

import (
	// "encoding/json"
	"github.com/streadway/amqp"
	"log"
	"relay/reports/ip_reports/rabbitmq/dbWrite"
	"fmt"
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

const (
	host     = "localhost"
	port     = "5672"
	user     = "rabbit"
	password = "Maro123!"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("Error! %s: %s", msg, err)
	}
}

func StartConsumer() {
	fmt.Println("rabbitmq consumer has been started........")
	conn, err := amqp.Dial("amqp://" + user + ":" + password + "@" + host + ":" + port + "/")
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

	q, err := ch.QueueDeclare(
		"relay_metrics", // name
		false,           // durable
		false,           // delete when usused
		true,            // exclusive
		false,           // no-wait
		nil,             // arguments
	)

	failOnError(err, "Failed to declare a queue")

	err = ch.QueueBind(
		q.Name,       // queue name
		"#",          // routing key
		"relay_jmta", // exchange
		false,
		nil)

	failOnError(err, "Failed to bind a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto ack
		false,  // exclusive
		false,  // no local
		false,  // no wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		// var data RabbitMQData
		var d amqp.Delivery
		for d = range msgs {
			// write logic here
			// log.Println(string(d.Body))
			dbWrite.InsertData(d)
			// json.Unmarshal(d.Body, &data)
			// log.Println("Data:", data)
			// log.Printf(" [x] %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}

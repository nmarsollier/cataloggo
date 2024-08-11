package r_consume

import (
	"encoding/json"
	"log"

	"github.com/nmarsollier/cataloggo/service"
	"github.com/nmarsollier/cataloggo/tools/env"
	"github.com/streadway/amqp"
)

func consumeOrders() error {
	conn, err := amqp.Dial(env.Get().RabbitURL)
	if err != nil {
		return err
	}
	defer conn.Close()

	chn, err := conn.Channel()
	if err != nil {
		return err
	}
	defer chn.Close()

	err = chn.ExchangeDeclare(
		"catalog", // name
		"direct",  // type
		false,     // durable
		false,     // auto-deleted
		false,     // internal
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return err
	}

	queue, err := chn.QueueDeclare(
		"catalog", // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return err
	}

	err = chn.QueueBind(
		queue.Name, // queue name
		"catalog",  // routing key
		"catalog",  // exchange
		false,
		nil)
	if err != nil {
		return err
	}

	mgs, err := chn.Consume(
		queue.Name, // queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	if err != nil {
		return err
	}

	log.Print("RabbitMQ consumeOrdersChannel conectado")

	go func() {
		for d := range mgs {
			newMessage := &service.ConsumeArticleValidation{}
			body := d.Body
			log.Print(string(body))

			err = json.Unmarshal(body, newMessage)
			if err == nil {
				switch newMessage.Type {
				case "article-data":
					service.ProcessArticleData(newMessage)
				case "article-exist":
					service.ProcessArticleData(newMessage)
				}
			} else {
				log.Print(err.Error())
			}
		}
	}()

	log.Print("Closed connection: ", <-conn.NotifyClose(make(chan *amqp.Error)))

	return nil
}

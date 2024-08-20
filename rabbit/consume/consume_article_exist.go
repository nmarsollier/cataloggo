package consume

import (
	"encoding/json"

	"github.com/golang/glog"
	"github.com/nmarsollier/cataloggo/service"
	"github.com/nmarsollier/cataloggo/tools/env"
	"github.com/streadway/amqp"
)

//	@Summary		Mensage Rabbit article_exist/article_exist
//	@Description	Otros microservicios nos solicitan validar articulos en el catalogo.
//	@Tags			Rabbit
//	@Accept			json
//	@Produce		json
//	@Param			article_exist	body	service.ConsumeArticleExist	true	"Message para article_exist"
//	@Router			/rabbit/article_exist [get]
//
// Validar Artículos
func consumeArticleExist() error {
	conn, err := amqp.Dial(env.Get().RabbitURL)
	if err != nil {
		glog.Error(err)

		return err
	}
	defer conn.Close()

	chn, err := conn.Channel()
	if err != nil {
		glog.Error(err)

		return err
	}
	defer chn.Close()

	err = chn.ExchangeDeclare(
		"article_exist", // name
		"direct",        // type
		false,           // durable
		false,           // auto-deleted
		false,           // internal
		false,           // no-wait
		nil,             // arguments
	)
	if err != nil {
		glog.Error(err)

		return err
	}

	queue, err := chn.QueueDeclare(
		"catalog_article_exist", // name
		false,                   // durable
		false,                   // delete when unused
		false,                   // exclusive
		false,                   // no-wait
		nil,                     // arguments
	)
	if err != nil {
		glog.Error(err)

		return err
	}

	err = chn.QueueBind(
		queue.Name,      // queue name
		"article_exist", // routing key
		"article_exist", // exchange
		false,
		nil)
	if err != nil {
		glog.Error(err)

		return err
	}

	mgs, err := chn.Consume(
		queue.Name, // queue
		"",         // consumer
		false,      // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	if err != nil {
		glog.Error(err)

		return err
	}

	glog.Info("RabbitMQ consumeOrdersChannel conectado")

	go func() {
		for d := range mgs {
			body := d.Body
			glog.Info("Incomming article_exist :", string(body))

			newMessage := &service.ConsumeArticleExist{}
			err = json.Unmarshal(body, newMessage)
			if err == nil {
				service.ProcessArticleData(newMessage)

				if err := d.Ack(false); err != nil {
					glog.Info("Failed ACK article_exist :", string(body), err)
				} else {
					glog.Info("Consumed article_exist :", string(body))
				}
			} else {
				glog.Error(err)
			}
		}
	}()

	glog.Info("Closed connection: ", <-conn.NotifyClose(make(chan *amqp.Error)))

	return nil
}

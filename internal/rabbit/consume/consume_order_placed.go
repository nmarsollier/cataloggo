package consume

import (
	"encoding/json"

	"github.com/nmarsollier/cataloggo/internal/env"
	"github.com/nmarsollier/cataloggo/internal/services"
	"github.com/nmarsollier/commongo/log"
	"github.com/streadway/amqp"
)

type OrderPlacedConsumer interface {
	ConsumeOrderPlaced() error
}

func NewOrderPlacedConsumer(fluentUrl string, rabbitUrl string, service services.CatalogService) OrderPlacedConsumer {
	logger := log.Get(fluentUrl, "cataloggo").
		WithField(log.LOG_FIELD_CONTROLLER, "Rabbit").
		WithField(log.LOG_FIELD_RABBIT_EXCHANGE, "order_placed").
		WithField(log.LOG_FIELD_RABBIT_QUEUE, "catalog_order_placed").
		WithField(log.LOG_FIELD_RABBIT_ACTION, "Consume")

	return &orderPlacedConsumer{
		logger:    logger,
		rabbitUrl: rabbitUrl,
		service:   service,
	}
}

type orderPlacedConsumer struct {
	logger    log.LogRusEntry
	rabbitUrl string
	service   services.CatalogService
}

//	@Summary		Mensage Rabbit order_placed/order_placed
//	@Description	Cuando se recibe el mensage order_placed damos de baja al stock para reservar los articulos. Queda pendiente enviar mensaje confirmando la operacion al MS de Orders.
//	@Tags			Rabbit
//	@Accept			json
//	@Produce		json
//	@Param			order_placed	body	services.ConsumeOrderPlacedMessage	true	"Message order_placed"
//	@Router			/rabbit/order_placed [get]
//
// Consume Order Placed
func (r *orderPlacedConsumer) ConsumeOrderPlaced() error {
	conn, err := amqp.Dial(env.Get().RabbitURL)
	if err != nil {
		return err
	}
	defer conn.Close()

	chn, err := conn.Channel()
	if err != nil {
		r.logger.Error(err)
		return err
	}
	defer chn.Close()

	err = chn.ExchangeDeclare(
		"order_placed", // name
		"fanout",       // type
		false,          // durable
		false,          // auto-deleted
		false,          // internal
		false,          // no-wait
		nil,            // arguments
	)
	if err != nil {
		r.logger.Error(err)
		return err
	}

	queue, err := chn.QueueDeclare(
		"catalog_order_placed", // name
		false,                  // durable
		false,                  // delete when unused
		false,                  // exclusive
		false,                  // no-wait
		nil,                    // arguments
	)
	if err != nil {
		r.logger.Error(err)
		return err
	}

	err = chn.QueueBind(
		queue.Name,     // queue name
		"",             // routing key
		"order_placed", // exchange
		false,
		nil)
	if err != nil {
		r.logger.Error(err)
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
		r.logger.Error(err)
		return err
	}

	r.logger.Info("RabbitMQ consumeOrderPlaced conectado")

	go func() {
		for d := range mgs {
			body := d.Body

			articleMessage := &services.ConsumeOrderPlaced{}
			err = json.Unmarshal(body, articleMessage)
			if err == nil {
				r.logger.Info("Incoming order_placed :", string(body))

				r.service.ProcessOrderPlaced(articleMessage)

				if err := d.Ack(false); err != nil {
					r.logger.Info("Failed ACK order_placed :", string(body), err)
				} else {
					r.logger.Info("Consumed order_placed :", string(body))
				}
			} else {
				r.logger.Error(err)
			}
		}
	}()

	r.logger.Info("Closed connection: ", <-conn.NotifyClose(make(chan *amqp.Error)))

	return nil
}

package consume

import (
	"encoding/json"

	"github.com/nmarsollier/cataloggo/log"
	"github.com/nmarsollier/cataloggo/service"
	"github.com/nmarsollier/cataloggo/tools/env"
	uuid "github.com/satori/go.uuid"
	"github.com/streadway/amqp"
)

//	@Summary		Mensage Rabbit order_placed/order_placed
//	@Description	Cuando se recibe el mensage order_placed damos de baja al stock para reservar los articulos. Queda pendiente enviar mensaje confirmando la operacion al MS de Orders.
//	@Tags			Rabbit
//	@Accept			json
//	@Produce		json
//	@Param			order_placed	body	service.ConsumeOrderPlacedMessage	true	"Message order_placed"
//	@Router			/rabbit/order_placed [get]
//
// Consume Order Placed
func consumeOrderPlaced() error {
	logger := log.Get().
		WithField(log.LOG_FIELD_CONTROLLER, "Rabbit").
		WithField(log.LOG_FIELD_RABBIT_EXCHANGE, "order_placed").
		WithField(log.LOG_FIELD_RABBIT_QUEUE, "catalog_order_placed").
		WithField(log.LOG_FIELD_RABBIT_ACTION, "Consume")

	conn, err := amqp.Dial(env.Get().RabbitURL)
	if err != nil {
		return err
	}
	defer conn.Close()

	chn, err := conn.Channel()
	if err != nil {
		logger.Error(err)
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
		logger.Error(err)
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
		logger.Error(err)
		return err
	}

	err = chn.QueueBind(
		queue.Name,     // queue name
		"",             // routing key
		"order_placed", // exchange
		false,
		nil)
	if err != nil {
		logger.Error(err)
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
		logger.Error(err)
		return err
	}

	logger.Info("RabbitMQ consumeOrderPlaced conectado")

	go func() {
		for d := range mgs {
			body := d.Body

			articleMessage := &service.ConsumeOrderPlaced{}
			err = json.Unmarshal(body, articleMessage)
			if err == nil {
				l := logger.WithField(log.LOG_FIELD_CORRELATION_ID, getConsumeOrderPlacedCorrelationId(articleMessage))
				l.Info("Incoming order_placed :", string(body))

				service.ProcessOrderPlaced(articleMessage, l)

				if err := d.Ack(false); err != nil {
					l.Info("Failed ACK order_placed :", string(body), err)
				} else {
					l.Info("Consumed order_placed :", string(body))
				}
			} else {
				logger.Error(err)
			}
		}
	}()

	logger.Info("Closed connection: ", <-conn.NotifyClose(make(chan *amqp.Error)))

	return nil
}

func getConsumeOrderPlacedCorrelationId(c *service.ConsumeOrderPlaced) string {
	value := c.CorrelationId

	if len(value) == 0 {
		value = uuid.NewV4().String()
	}

	return value
}

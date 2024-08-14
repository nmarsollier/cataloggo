package r_consume

import (
	"encoding/json"

	"github.com/golang/glog"
	"github.com/nmarsollier/cataloggo/service"
	"github.com/nmarsollier/cataloggo/tools/env"
	"github.com/streadway/amqp"
)

//	@Summary		Mensage Rabbit order/order-placed
//	@Description	Cuando se recibe el mensage order-placed damos de baja al stock para reservar los articulos. Queda pendiente enviar mensaje confirmando la operacion al MS de Orders.
//	@Tags			Rabbit
//	@Accept			json
//	@Produce		json
//	@Param			article-data	body	service.ConsumeOrderPlacedMessage	true	"Message para Type = article-data"
//	@Router			/rabbit/order-placed [get]
//
// Consume Order Placed
func consumeOrderPlaced() error {
	conn, err := amqp.Dial(env.Get().RabbitURL)
	if err != nil {
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
		"order_placed", // name
		"fanout",       // type
		false,          // durable
		false,          // auto-deleted
		false,          // internal
		false,          // no-wait
		nil,            // arguments
	)
	if err != nil {
		glog.Error(err)
		return err
	}

	queue, err := chn.QueueDeclare(
		"cart_order_placed", // name
		false,               // durable
		false,               // delete when unused
		false,               // exclusive
		false,               // no-wait
		nil,                 // arguments
	)
	if err != nil {
		glog.Error(err)
		return err
	}

	err = chn.QueueBind(
		queue.Name,          // queue name
		"cart_order_placed", // routing key
		"order_placed",      // exchange
		false,
		nil)
	if err != nil {
		glog.Error(err)
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
		glog.Error(err)
		return err
	}

	glog.Info("RabbitMQ consumeOrderPlaced conectado")

	go func() {
		for d := range mgs {
			newMessage := &ConsumeMessage{}
			body := d.Body
			glog.Info("Rannit Consumed : ", string(body))

			err = json.Unmarshal(body, newMessage)
			if err == nil {
				switch newMessage.Type {
				case "order-placed":
					articleMessage := &service.ConsumeOrderPlaced{}
					if err := json.Unmarshal(body, articleMessage); err != nil {
						glog.Error(err)
						return
					}

					service.ProcessOrderPlaced(articleMessage)
				}
			} else {
				glog.Error(err)
			}
		}
	}()

	glog.Info("Closed connection: ", <-conn.NotifyClose(make(chan *amqp.Error)))

	return nil
}

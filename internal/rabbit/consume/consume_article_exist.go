package consume

import (
	"encoding/json"

	"github.com/nmarsollier/cataloggo/internal/rabbit/rschema"
	"github.com/nmarsollier/cataloggo/internal/services"
	"github.com/nmarsollier/commongo/log"

	uuid "github.com/satori/go.uuid"
	"github.com/streadway/amqp"
)

type ArticleExistConsumer interface {
	ConsumeArticleExist() error
}

func NewArticleExistConsumer(fluentUrl string, rabbitUrl string, service services.CatalogService) ArticleExistConsumer {
	log := log.Get(fluentUrl, "cataloggo").
		WithField(log.LOG_FIELD_CONTROLLER, "Rabbit").
		WithField(log.LOG_FIELD_RABBIT_EXCHANGE, "article_exist").
		WithField(log.LOG_FIELD_RABBIT_QUEUE, "catalog_article_exist").
		WithField(log.LOG_FIELD_RABBIT_ACTION, "Consume")

	return &articleExistConsumer{
		log:       log,
		rabbitUrl: rabbitUrl,
		service:   service,
	}
}

type articleExistConsumer struct {
	log       log.LogRusEntry
	rabbitUrl string
	service   services.CatalogService
}

//	@Summary		Mensage Rabbit article_exist/article_exist
//	@Description	Otros microservicios nos solicitan validar articulos en el catalogo.
//	@Tags			Rabbit
//	@Accept			json
//	@Produce		json
//	@Param			article_exist	body	rschema.ConsumeArticleExist	true	"Message para article_exist"
//	@Router			/rabbit/article_exist [get]
//
// Validar Artículos
func (r *articleExistConsumer) ConsumeArticleExist() error {
	conn, err := amqp.Dial(r.rabbitUrl)
	if err != nil {
		r.log.Error(err)

		return err
	}
	defer conn.Close()

	chn, err := conn.Channel()
	if err != nil {
		r.log.Error(err)

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
		r.log.Error(err)

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
		r.log.Error(err)

		return err
	}

	err = chn.QueueBind(
		queue.Name,      // queue name
		"article_exist", // routing key
		"article_exist", // exchange
		false,
		nil)
	if err != nil {
		r.log.Error(err)

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
		r.log.Error(err)

		return err
	}

	r.log.Info("RabbitMQ consumeOrdersChannel conectado")

	go func() {
		for d := range mgs {
			body := d.Body

			newMessage := &rschema.ConsumeArticleExist{}
			err = json.Unmarshal(body, newMessage)
			if err == nil {
				l := r.log.WithField(log.LOG_FIELD_CORRELATION_ID, getConsumeArticleExistCorrelationId(newMessage))
				l.Info("Incoming article_exist :", string(body))
				r.service.ProcessArticleData(newMessage)

				if err := d.Ack(false); err != nil {
					l.Info("Failed ACK article_exist :", string(body), err)
				} else {
					l.Info("Consumed article_exist :", string(body))
				}
			} else {
				r.log.Error(err)
			}
		}
	}()

	r.log.Info("Closed connection: ", <-conn.NotifyClose(make(chan *amqp.Error)))

	return nil
}

func getConsumeArticleExistCorrelationId(c *rschema.ConsumeArticleExist) string {
	value := c.CorrelationId

	if len(value) == 0 {
		value = uuid.NewV4().String()
	}

	return value
}

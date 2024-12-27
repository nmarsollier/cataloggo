package emit

import (
	"encoding/json"
	"errors"

	"github.com/nmarsollier/cataloggo/internal/engine/env"
	"github.com/nmarsollier/cataloggo/internal/engine/log"
	"github.com/nmarsollier/cataloggo/internal/rabbit/rschema"
	"github.com/streadway/amqp"
)

type RabbitEmitter interface {
	EmitArticleExist(exchange string, routingKey string, message *rschema.ArticleExistMessage, deps ...interface{}) error
}

func NewRabbitEmitter(fluentUrl string) RabbitEmitter {
	logger := log.Get(fluentUrl).
		WithField(log.LOG_FIELD_CONTROLLER, "Rabbit").
		WithField(log.LOG_FIELD_RABBIT_ACTION, "Emit").
		WithField(log.LOG_FIELD_RABBIT_EXCHANGE, "article_exist")
	return &rabbitEmitter{
		log: logger,
	}
}

type rabbitEmitter struct {
	log log.LogRusEntry
}

// ErrChannelNotInitialized Rabbit channel could not be initialized
var ErrChannelNotInitialized = errors.New("channel not initialized")

func (e *rabbitEmitter) getChannel() (*amqp.Channel, error) {
	conn, err := amqp.Dial(env.Get().RabbitURL)
	if err != nil {
		e.log.Error(err)
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		e.log.Error(err)
		return nil, err
	}
	if ch == nil {
		e.log.Error(err)
		return nil, ErrChannelNotInitialized
	}
	return ch, nil
}

// Emite respuestas de article_exist
//
//	@Summary		Mensage Rabbit article_exist
//	@Description	Emite respuestas de article_exist
//	@Tags			Rabbit
//	@Accept			json
//	@Produce		json
//	@Param			body	body	rschema.SendArticleExist	true	"Estructura general del mensage"
//	@Router			/rabbit/article_exist [put]
func (e *rabbitEmitter) EmitArticleExist(exchange string, routingKey string, message *rschema.ArticleExistMessage, deps ...interface{}) error {

	corrId, _ := e.log.Data()[log.LOG_FIELD_CORRELATION_ID].(string)
	data := &rschema.SendArticleExist{
		Message:       *message,
		CorrelationId: corrId,
	}

	chn, err := e.getChannel()
	if err != nil {
		chn = nil
		e.log.Error(err)
		return err
	}

	err = chn.ExchangeDeclare(
		exchange, // name
		"direct", // type
		false,    // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		chn = nil
		e.log.Error(err)
		return err
	}

	body, err := json.Marshal(data)
	if err != nil {
		e.log.Error(err)
		return err
	}

	err = chn.Publish(
		exchange,   // exchange
		routingKey, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			Body: []byte(body),
		})
	if err != nil {
		chn = nil
		e.log.Error(err)
		return err
	}

	e.log.Info(string(body))

	return nil
}

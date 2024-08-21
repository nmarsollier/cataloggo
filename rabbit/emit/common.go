package emit

import (
	"encoding/json"
	"errors"

	"github.com/nmarsollier/cataloggo/log"
	"github.com/nmarsollier/cataloggo/rabbit/rschema"
	"github.com/nmarsollier/cataloggo/tools/env"
	"github.com/streadway/amqp"
)

// ErrChannelNotInitialized Rabbit channel could not be initialized
var ErrChannelNotInitialized = errors.New("channel not initialized")

func getChannel(ctx ...interface{}) (*amqp.Channel, error) {
	conn, err := amqp.Dial(env.Get().RabbitURL)
	if err != nil {
		log.Get(ctx...).Error(err)
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Get(ctx...).Error(err)
		return nil, err
	}
	if ch == nil {
		log.Get(ctx...).Error(err)
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
func EmitArticleExist(exchange string, routingKey string, message *rschema.ArticleExistMessage, ctx ...interface{}) error {

	logger := log.Get(ctx...).
		WithField(log.LOG_FIELD_CONTOROLLER, "Rabbit").
		WithField(log.LOG_FIELD_RABBIT_ACTION, "Emit").
		WithField(log.LOG_FIELD_RABBIT_EXCHANGE, "article_exist")
	corrId, _ := logger.Data[log.LOG_FIELD_CORRELATION_ID].(string)
	data := &rschema.SendArticleExist{
		Message:       *message,
		CorrelationId: corrId,
	}

	chn, err := getChannel(ctx...)
	if err != nil {
		chn = nil
		logger.Error(err)
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
		logger.Error(err)
		return err
	}

	body, err := json.Marshal(data)
	if err != nil {
		logger.Error(err)
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
		logger.Error(err)
		return err
	}

	logger.Info(string(body))

	return nil
}

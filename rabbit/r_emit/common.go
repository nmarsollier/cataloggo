package r_emit

import (
	"encoding/json"
	"log"

	"github.com/nmarsollier/cataloggo/tools"
	"github.com/nmarsollier/cataloggo/tools/env"
	"github.com/nmarsollier/cataloggo/tools/errors"
	"github.com/streadway/amqp"
)

// ErrChannelNotInitialized Rabbit channel could not be initialized
var ErrChannelNotInitialized = errors.NewCustom(400, "Channel not initialized")

func getChannel() (*amqp.Channel, error) {
	conn, err := amqp.Dial(env.Get().RabbitURL)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	if ch == nil {
		return nil, ErrChannelNotInitialized
	}
	return ch, nil
}

// Emite respuestas de article-data or article-exist
//
//	@Summary		Mensage Rabbit
//	@Description	Emite respuestas de article-data or article-exist
//	@Tags			Rabbit
//	@Accept			json
//	@Produce		json
//	@Param			body	body	service.EmitArticleValidation	true	"Estructura general del mensage"
//	@Router			/rabbit/article-data [put]
func EmitDirect(exchange string, queue string, data interface{}) error {

	chn, err := getChannel()
	if err != nil {
		chn = nil
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
		return err
	}

	body, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = chn.Publish(
		exchange, // exchange
		queue,    // routing key
		false,    // mandatory
		false,    // immediate
		amqp.Publishing{
			Body: []byte(body),
		})
	if err != nil {
		chn = nil
		return err
	}

	log.Output(1, "Rabbit enviado "+tools.ToJson(body))

	return nil
}

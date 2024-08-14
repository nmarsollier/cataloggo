package emit

import (
	"encoding/json"
	"errors"

	"github.com/golang/glog"
	"github.com/nmarsollier/cataloggo/tools/env"
	"github.com/streadway/amqp"
)

// ErrChannelNotInitialized Rabbit channel could not be initialized
var ErrChannelNotInitialized = errors.New("channel not initialized")

func getChannel() (*amqp.Channel, error) {
	conn, err := amqp.Dial(env.Get().RabbitURL)
	if err != nil {
		glog.Error(err)
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		glog.Error(err)
		return nil, err
	}
	if ch == nil {
		glog.Error(err)
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
		glog.Error(err)
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
		glog.Error(err)
		return err
	}

	body, err := json.Marshal(data)
	if err != nil {
		glog.Error(err)
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
		glog.Error(err)
		return err
	}

	glog.Info("Rabbit Sent : ", body)

	return nil
}

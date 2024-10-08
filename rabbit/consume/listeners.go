package consume

import (
	"time"

	"github.com/nmarsollier/cataloggo/tools/log"
)

func Init() {
	logger := log.Get().
		WithField(log.LOG_FIELD_CONTROLLER, "Rabbit").
		WithField(log.LOG_FIELD_RABBIT_ACTION, "Init")

	go func() {
		for {
			err := consumeArticleExist()
			if err != nil {
				logger.Error(err)
			}
			logger.Info("RabbitMQ consumeOrders conectando en 5 segundos.")
			time.Sleep(5 * time.Second)
		}
	}()
	go func() {
		for {
			err := consumeOrderPlaced()
			if err != nil {
				logger.Error(err)
			}
			logger.Info("RabbitMQ consumeOrderPlaced conectando en 5 segundos.")
			time.Sleep(5 * time.Second)
		}
	}()
	go func() {
		for {
			err := consumeLogout()
			if err != nil {
				logger.Error(err)
			}
			logger.Info("RabbitMQ listenLogout conectando en 5 segundos.")
			time.Sleep(5 * time.Second)
		}
	}()
}

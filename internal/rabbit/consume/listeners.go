package consume

import (
	"time"

	"github.com/nmarsollier/cataloggo/internal/engine/log"
)

func Init(logger log.LogRusEntry, articleExistConsumer ArticleExistConsumer, logoutConsumer LogoutConsumer, orderPlacedConsumer OrderPlacedConsumer) {
	logger.
		WithField(log.LOG_FIELD_CONTROLLER, "Rabbit").
		WithField(log.LOG_FIELD_RABBIT_ACTION, "Init")

	go func() {
		for {
			err := articleExistConsumer.ConsumeArticleExist()
			if err != nil {
				logger.Error(err)
			}
			logger.Info("RabbitMQ consumeOrders conectando en 5 segundos.")
			time.Sleep(5 * time.Second)
		}
	}()
	go func() {
		for {
			err := orderPlacedConsumer.ConsumeOrderPlaced()
			if err != nil {
				logger.Error(err)
			}
			logger.Info("RabbitMQ consumeOrderPlaced conectando en 5 segundos.")
			time.Sleep(5 * time.Second)
		}
	}()
	go func() {
		for {
			err := logoutConsumer.ConsumeLogout()
			if err != nil {
				logger.Error(err)
			}
			logger.Info("RabbitMQ listenLogout conectando en 5 segundos.")
			time.Sleep(5 * time.Second)
		}
	}()
}

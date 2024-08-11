package r_consume

import (
	"log"
	"time"
)

func Init() {
	go func() {
		for {
			err := consumeOrders()
			if err != nil {
				log.Print(err)
			}
			log.Print("RabbitMQ consumeOrders conectando en 5 segundos.")
			time.Sleep(5 * time.Second)
		}
	}()
	go func() {
		for {
			err := consumeOrderPlaced()
			if err != nil {
				log.Print(err)
			}
			log.Print("RabbitMQ consumeOrderPlaced conectando en 5 segundos.")
			time.Sleep(5 * time.Second)
		}
	}()
	go func() {
		for {
			err := consumeLogout()
			if err != nil {
				log.Print(err)
			}
			log.Print("RabbitMQ listenLogout conectando en 5 segundos.")
			time.Sleep(5 * time.Second)
		}
	}()
}

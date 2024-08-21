package service

import (
	"github.com/nmarsollier/cataloggo/article"
)

func ProcessOrderPlaced(data *ConsumeOrderPlaced, ctx ...interface{}) {
	for _, a := range data.Message.Articles {
		art, err := article.FindById(a.ArticleId, ctx...)
		if err == nil {
			article.DecrementStock(art.ID, a.Quantity, ctx...)
		}
	}
}

type ConsumeOrderPlaced struct {
	CorrelationId string `json:"correlation_id" example:"123123" `
	RoutingKey    string `json:"routing_key" example:"Remote RoutingKey to Reply"`
	Exchange      string `json:"exchange" example:"order-placed"`
	Message       *ConsumeOrderPlacedMessage
}

type ConsumeOrderPlacedMessage struct {
	OrderId  string                       `json:"orderId"`
	CartId   string                       `json:"cartId"`
	Articles []*ConsumeOrderPlacedArticle `json:"articles"`
}

type ConsumeOrderPlacedArticle struct {
	ArticleId string `json:"articleId"`
	Quantity  int    `json:"quantity"`
}

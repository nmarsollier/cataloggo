package service

import (
	"github.com/nmarsollier/cataloggo/article"
)

func ProcessOrderPlaced(data *ConsumeOrderPlaced) {
	for _, a := range data.Message.Articles {
		art, err := article.FindById(a.ArticleId)
		if err == nil {
			article.DecrementStock(art.ID, a.Quantity)
		}
	}
}

type ConsumeOrderPlaced struct {
	RoutingKey string `json:"routing_key" example:"Remote RoutingKey to Reply"`
	Exchange   string `json:"exchange" example:"order-placed"`
	Message    *ConsumeOrderPlacedMessage
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

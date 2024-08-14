package service

import (
	"github.com/golang/glog"
	"github.com/nmarsollier/cataloggo/article"
)

func ProcessOrderPlaced(data *ConsumeOrderPlaced) {
	for _, a := range data.Message.Articles {
		art, err := article.FindById(a.ArticleId)
		if err == nil {
			article.DecreaseStock(art.ID, art.Stock)
		}
	}

	// TODO: Enviar mensaje al MS de orders para confirmar que se dio de baja al stock.
	glog.Info("Order Placed processed : ", data)
}

type ConsumeOrderPlaced struct {
	Type     string `json:"type" example:"order-placed"`
	Queue    string `json:"queue" example:"order-placed"`
	Exchange string `json:"exchange" example:"order-placed"`
	Message  *ConsumeOrderPlacedMessage
}

type ConsumeOrderPlacedMessage struct {
	OrderId  string                       `json:"orderId"`
	CartId   int                          `json:"cartId"`
	Articles []*ConsumeOrderPlacedArticle `json:"articles"`
}

type ConsumeOrderPlacedArticle struct {
	ArticleId string `json:"articleId"`
	Quantity  int    `json:"quantity"`
}

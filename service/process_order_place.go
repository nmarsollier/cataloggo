package service

import (
	"github.com/golang/glog"
	"github.com/nmarsollier/cataloggo/article"
	"github.com/nmarsollier/cataloggo/tools"
)

// Consume Order Placed
//
//	@Summary		Mensage Rabbit order/order-placed
//	@Description	Antes de iniciar las operaciones se validan los artículos contra el catalogo.
//	@Tags			Rabbit
//	@Accept			json
//	@Produce		json
//	@Param			article-data	body	ConsumeOrderPlacedMessage	true	"Message para Type = article-data"
//
//	@Router			/rabbit/order-placed [get]
func ProcessOrderPlaced(data *ConsumeOrderPlaced) {
	for _, a := range data.Message.Articles {
		art, err := article.FindById(a.ArticleId)
		if err == nil {
			article.DecreaseStock(art.ID, art.Stock)
		}
	}

	glog.Info("Order Placed processed : ", tools.ToJson(data))
}

type ConsumeOrderPlaced struct {
	Type     string `json:"type"`
	Version  int    `json:"version"`
	Queue    string `json:"queue"`
	Exchange string `json:"exchange"`
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

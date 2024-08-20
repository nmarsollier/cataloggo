package service

import (
	"github.com/nmarsollier/cataloggo/article"
	"github.com/nmarsollier/cataloggo/rabbit/emit"
)

func ProcessArticleData(data *ConsumeArticleExist) {
	response := &SendArticleExist{
		Message: ArticleExistMessage{
			ArticleId:   data.Message.ArticleId,
			ReferenceId: data.Message.ReferenceId,
			Valid:       false,
		},
	}
	article, err := article.FindById(data.Message.ArticleId)
	if err != nil {
		emit.EmitDirect(data.Exchange, data.RoutingKey, response)
		return
	}

	response.Message = ArticleExistMessage{
		ArticleId:   data.Message.ArticleId,
		ReferenceId: data.Message.ReferenceId,
		Stock:       article.Stock,
		Price:       article.Price,
		Valid:       article.Enabled,
	}
	emit.EmitDirect(data.Exchange, data.RoutingKey, response)
}

type ArticleExistMessage struct {
	ArticleId   string  `json:"articleId" example:"ArticleId" `
	Price       float32 `json:"price"`
	ReferenceId string  `json:"referenceId" example:"Remote Reference Id"`
	Stock       int     `json:"stock"`
	Valid       bool    `json:"valid"`
}

type SendArticleExist struct {
	Message ArticleExistMessage `json:"message"`
}

type ConsumeArticleExist struct {
	RoutingKey string `json:"routing_key" example:"Remote RoutingKey to Reply"`
	Exchange   string `json:"exchange" example:"Remote Exchange to Reply"`
	Message    *ConsumeArticleExistMessage
}

type ConsumeArticleExistMessage struct {
	ReferenceId string `json:"referenceId" example:"Remote Reference Object Id"`
	ArticleId   string `json:"articleId" example:"ArticleId"`
}

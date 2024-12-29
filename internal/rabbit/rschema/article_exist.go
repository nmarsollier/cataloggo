package rschema

import "github.com/nmarsollier/commongo/rbt"

type ArticleExistPublisher = rbt.RabbitPublisher[*ArticleExistMessage]

type ArticleExistMessage struct {
	ArticleId   string  `json:"articleId" example:"ArticleId" `
	Price       float32 `json:"price"`
	ReferenceId string  `json:"referenceId" example:"Remote Reference Id"`
	Stock       int     `json:"stock"`
	Valid       bool    `json:"valid"`
}

type SendArticleExist struct {
	CorrelationId string              `json:"correlation_id" example:"123123" `
	Message       ArticleExistMessage `json:"message"`
}

type ConsumeArticleExist struct {
	CorrelationId string `json:"correlation_id" example:"123123" `
	RoutingKey    string `json:"routing_key" example:"Remote RoutingKey to Reply"`
	Exchange      string `json:"exchange" example:"Remote Exchange to Reply"`
	Message       *ConsumeArticleExistMessage
}

type ConsumeArticleExistMessage struct {
	ReferenceId string `json:"referenceId" example:"Remote Reference Object Id"`
	ArticleId   string `json:"articleId" example:"ArticleId"`
}

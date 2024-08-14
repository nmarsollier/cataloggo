package service

import (
	"github.com/golang/glog"
	"github.com/nmarsollier/cataloggo/article"
	"github.com/nmarsollier/cataloggo/rabbit/emit"
)

func ProcessArticleData(data *ConsumeArticleValidation) {
	response := &SendValidationMessage{
		Type:     data.Type,
		Exchange: data.Exchange,
		Queue:    data.Queue,
		Message: EmitArticleValidation{
			ArticleId:   data.Message.ArticleId,
			ReferenceId: data.Message.ReferenceId,
			Valid:       false,
		},
	}
	article, err := article.FindById(data.Message.ArticleId)
	if err != nil {
		emit.EmitDirect(data.Exchange, data.Queue, response)
		return
	}

	response.Message = EmitArticleValidation{
		ArticleId:   data.Message.ArticleId,
		ReferenceId: data.Message.ReferenceId,
		Stock:       article.Stock,
		Price:       article.Price,
		Valid:       article.Enabled,
	}
	emit.EmitDirect(data.Exchange, data.Queue, response)

	glog.Info("Article validation completed : ", data)
}

type EmitArticleValidation struct {
	ArticleId   string  `json:"articleId" example:"ArticleId" `
	Price       float32 `json:"price"`
	ReferenceId string  `json:"referenceId" example:"Remote Reference Id"`
	Stock       int     `json:"stock"`
	Valid       bool    `json:"valid"`
}

type ConsumeArticleValidation struct {
	Type     string `json:"type" example:"article-data" `
	Queue    string `json:"queue" example:"Remote Queue to Reply" `
	Exchange string `json:"exchange" example:"Remote Exchange to Reply"`
	Message  *ConsumeArticleValidationMessage
}

type SendValidationMessage struct {
	Type     string      `json:"type" example:"article-exist"`
	Exchange string      `json:"exchange" example:"cart"`
	Queue    string      `json:"queue" example:"cart"`
	Message  interface{} `json:"message"`
}

type ConsumeArticleValidationMessage struct {
	ReferenceId string `json:"referenceId" example:"Remote Reference Object Id"`

	ArticleId string `json:"articleId" example:"ArticleId"`
}

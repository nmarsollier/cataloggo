package services

import (
	"github.com/nmarsollier/cataloggo/article"
	"github.com/nmarsollier/cataloggo/rabbit/emit"
	"github.com/nmarsollier/cataloggo/rabbit/rschema"
)

func ProcessArticleData(data *rschema.ConsumeArticleExist, ctx ...interface{}) {
	article, err := article.FindById(data.Message.ArticleId, ctx...)
	if err != nil {
		emit.EmitArticleExist(data.Exchange, data.RoutingKey, &rschema.ArticleExistMessage{
			ArticleId:   data.Message.ArticleId,
			ReferenceId: data.Message.ReferenceId,
			Valid:       false,
		}, ctx...)
		return
	}

	emit.EmitArticleExist(data.Exchange, data.RoutingKey, &rschema.ArticleExistMessage{
		ArticleId:   data.Message.ArticleId,
		ReferenceId: data.Message.ReferenceId,
		Stock:       article.Stock,
		Price:       article.Price,
		Valid:       article.Enabled,
	}, ctx...)
}

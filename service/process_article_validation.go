package service

import (
	"github.com/golang/glog"
	"github.com/nmarsollier/cataloggo/article"
	"github.com/nmarsollier/cataloggo/rabbit/r_emit"
	"github.com/nmarsollier/cataloggo/tools"
)

// @Summary		Mensage Rabbit article-data o article-exist
// @Description	Antes de iniciar las operaciones se validan los artículos contra el catalogo.
// @Tags			Rabbit
// @Accept			json
// @Produce		json
// @Param			article-data	body	ConsumeArticleValidation	true	"Message para Type = article-data"
// @Router			/rabbit/article-data [get]
//
// Validar Artículos
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
		r_emit.EmitDirect(data.Exchange, data.Queue, response)
		return
	}

	response.Message = EmitArticleValidation{
		ArticleId:   data.Message.ArticleId,
		ReferenceId: data.Message.ReferenceId,
		Stock:       article.Stock,
		Price:       article.Price,
		Valid:       article.Enabled,
	}
	r_emit.EmitDirect(data.Exchange, data.Queue, response)

	glog.Info("Article validation completed : ", tools.ToJson(data))
}

type EmitArticleValidation struct {
	ArticleId   string  `json:"articleId"`
	Price       float32 `json:"price"`
	ReferenceId string  `json:"referenceId"`
	Stock       int     `json:"stock"`
	Valid       bool    `json:"valid"`
}

type ConsumeArticleValidation struct {
	Type     string `json:"type"`
	Version  int    `json:"version"`
	Queue    string `json:"queue"`
	Exchange string `json:"exchange"`
	Message  *ConsumeArticleValidationMessage
}

type SendValidationMessage struct {
	Type     string      `json:"type"`
	Exchange string      `json:"exchange"`
	Queue    string      `json:"queue"`
	Message  interface{} `json:"message"`
}

type ConsumeArticleValidationMessage struct {
	ArticleId   string `json:"articleId"`
	ReferenceId string `json:"referenceId"`
}

package services

import (
	"github.com/nmarsollier/cataloggo/internal/article"
	"github.com/nmarsollier/cataloggo/internal/rabbit/emit"
	"github.com/nmarsollier/cataloggo/internal/rabbit/rschema"
)

type CatalogService interface {
	ProcessArticleData(data *rschema.ConsumeArticleExist)
	ProcessOrderPlaced(data *ConsumeOrderPlaced)
}

func NewCatalogService(catalogService article.ArticleService, emit emit.RabbitEmitter) CatalogService {
	return &catService{
		catalogService: catalogService,
		emit:           emit,
	}
}

type catService struct {
	catalogService article.ArticleService
	emit           emit.RabbitEmitter
}

func (s *catService) ProcessArticleData(data *rschema.ConsumeArticleExist) {
	article, err := s.catalogService.FindById(data.Message.ArticleId)
	if err != nil {
		s.emit.EmitArticleExist(data.Exchange, data.RoutingKey, &rschema.ArticleExistMessage{
			ArticleId:   data.Message.ArticleId,
			ReferenceId: data.Message.ReferenceId,
			Valid:       false,
		})
		return
	}

	s.emit.EmitArticleExist(data.Exchange, data.RoutingKey, &rschema.ArticleExistMessage{
		ArticleId:   data.Message.ArticleId,
		ReferenceId: data.Message.ReferenceId,
		Stock:       article.Stock,
		Price:       article.Price,
		Valid:       article.Enabled,
	})
}

func (s *catService) ProcessOrderPlaced(data *ConsumeOrderPlaced) {
	for _, a := range data.Message.Articles {
		art, err := s.catalogService.FindById(a.ArticleId)
		if err == nil {
			s.catalogService.DecrementStock(art.ID, a.Quantity)
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

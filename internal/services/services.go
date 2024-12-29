package services

import (
	"github.com/nmarsollier/cataloggo/internal/article"
	"github.com/nmarsollier/cataloggo/internal/rabbit/rschema"
	"github.com/nmarsollier/commongo/log"
	uuid "github.com/satori/go.uuid"
)

type CatalogService interface {
	ProcessArticleData(data *rschema.ConsumeArticleExist)
	ProcessOrderPlaced(data *ConsumeOrderPlaced)
}

func NewCatalogService(log log.LogRusEntry, catalogService article.ArticleService, publisher rschema.ArticleExistPublisher) CatalogService {
	return &catService{
		log:            log,
		catalogService: catalogService,
		emit:           publisher,
	}
}

type catService struct {
	log            log.LogRusEntry
	catalogService article.ArticleService
	emit           rschema.ArticleExistPublisher
}

func (s *catService) ProcessArticleData(data *rschema.ConsumeArticleExist) {
	article, err := s.catalogService.FindById(data.Message.ArticleId)
	if err != nil {
		s.emit.PublishTo(
			getConsumeArticleExistCorrelationId(data),
			data.Exchange,
			data.RoutingKey,
			&rschema.ArticleExistMessage{
				ArticleId:   data.Message.ArticleId,
				ReferenceId: data.Message.ReferenceId,
				Valid:       false,
			},
		)
		return
	}

	s.emit.PublishTo(
		getConsumeArticleExistCorrelationId(data),
		data.Exchange,
		data.RoutingKey,
		&rschema.ArticleExistMessage{
			ArticleId:   data.Message.ArticleId,
			ReferenceId: data.Message.ReferenceId,
			Stock:       article.Stock,
			Price:       article.Price,
			Valid:       article.Enabled,
		},
	)
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

func getConsumeArticleExistCorrelationId(c *rschema.ConsumeArticleExist) string {
	value := c.CorrelationId

	if len(value) == 0 {
		value = uuid.NewV4().String()
	}

	return value
}

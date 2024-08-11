package article

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FindById(id string) (*ArticleData, error) {

	article, err := findById(id)
	if err != nil {
		return nil, err
	}

	return toArticleData(article), nil
}

func FindByCriteria(criteria string) ([]*ArticleData, error) {
	articles, err := findByCriteria(criteria)
	if err != nil {
		return nil, err
	}

	result := []*ArticleData{}
	for _, a := range articles {
		result = append(result, toArticleData(a))
	}

	return result, nil
}

func CreateArticle(articleData *NewArticleData) (*ArticleData, error) {
	article, err := insert(&Article{
		ID: primitive.NewObjectID(),
		Description: Description{
			Name:        articleData.Name,
			Description: articleData.Description,
			Image:       articleData.Image,
		},
		Price:   articleData.Price,
		Stock:   articleData.Stock,
		Enabled: true,
		Created: time.Now(),
		Updated: time.Now(),
	})

	if err != nil {
		return nil, err
	}
	return toArticleData(article), nil
}

func UpdateArticle(articleId string, articleData *NewArticleData) error {
	err := updateDescription(articleId, Description{
		Name:        articleData.Name,
		Description: articleData.Description,
		Image:       articleData.Image,
	})

	if err != nil {
		return err
	}

	err = updateStock(articleId, articleData.Stock)

	if err != nil {
		return err
	}

	err = updatePrice(articleId, articleData.Price)

	if err != nil {
		return err
	}

	return nil
}

func toArticleData(article *Article) *ArticleData {
	return &ArticleData{
		ID:          article.ID,
		Name:        article.Description.Name,
		Description: article.Description.Description,
		Image:       article.Description.Image,
		Price:       article.Price,
		Stock:       article.Stock,
		Enabled:     article.Enabled,
	}
}

type ArticleData struct {
	ID          primitive.ObjectID `json:"_id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Image       string             `json:"image"`
	Price       float32            `json:"price"`
	Stock       int                `json:"stock"`
	Enabled     bool               `json:"enabled"`
}

type NewArticleData struct {
	Name string `json:"name" validate:"required,min=1,max=100"`

	Description string `json:"description" validate:"required,min=1,max=256"`
	Image       string `json:"image" validate:"max=100"`

	Price float32 `json:"price" validate:"required,min=1"`
	Stock int     `json:"stock" validate:"required,min=1"`
}

package article

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateArticle(articleData *UpdateArticleData, deps ...interface{}) (*ArticleData, error) {
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
	}, deps...)

	if err != nil {
		return nil, err
	}
	return toArticleData(article), nil
}

package article

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

func CreateArticle(articleData *UpdateArticleData, deps ...interface{}) (*ArticleData, error) {
	article := &Article{
		ID:          uuid.NewV4().String(),
		Name:        articleData.Name,
		Description: articleData.Description,
		Image:       articleData.Image,
		Price:       articleData.Price,
		Stock:       articleData.Stock,
		Enabled:     true,
		Created:     time.Now(),
		Updated:     time.Now(),
	}

	err := insert(article, deps...)

	if err != nil {
		return nil, err
	}

	return toArticleData(article), nil
}

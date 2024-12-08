package article

type ArticleData struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Price       float32 `json:"price"`
	Stock       int     `json:"stock"`
	Enabled     bool    `json:"enabled"`
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

type UpdateArticleData struct {
	Name string `dynamodbav:"name"  json:"name" validate:"required,min=1,max=100"`

	Description string `dynamodbav:"description" json:"description" validate:"required,min=1,max=256"`
	Image       string `dynamodbav:"image" json:"image" validate:"max=100"`

	Price float32 `dynamodbav:"price" json:"price" validate:"required,min=1"`
	Stock int     `dynamodbav:"stock" json:"stock" validate:"required,min=1"`
}

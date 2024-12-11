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
		Name:        article.Name,
		Description: article.Description,
		Image:       article.Image,
		Price:       article.Price,
		Stock:       article.Stock,
		Enabled:     article.Enabled,
	}
}

type UpdateArticleData struct {
	Name string `json:"name" validate:"required,min=1,max=100"`

	Description string `json:"description" validate:"required,min=1,max=256"`
	Image       string `json:"image" validate:"max=100"`

	Price float32 ` json:"price" validate:"required,min=1"`
	Stock int     ` json:"stock" validate:"required,min=1"`
}

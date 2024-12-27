package resolvers

import (
	"context"

	"github.com/nmarsollier/cataloggo/internal/article"
	"github.com/nmarsollier/cataloggo/internal/graph/model"
	"github.com/nmarsollier/cataloggo/internal/graph/tools"
)

func UpdateArticle(ctx context.Context, input model.UpdateArticle) (bool, error) {
	_, err := tools.ValidateAdmin(ctx)
	if err != nil {
		return false, err
	}

	env := tools.GqlDi(ctx)

	err = env.ArticleService().UpdateArticle(input.ID, &article.UpdateArticleData{
		Name:        input.Name,
		Description: input.Description,
		Image:       input.Image,
		Price:       float32(input.Price),
		Stock:       input.Stock,
	})

	if err != nil {
		return false, err
	}

	return true, nil
}
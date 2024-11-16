package resolvers

import (
	"context"

	"github.com/nmarsollier/cataloggo/article"
	"github.com/nmarsollier/cataloggo/graph/model"
	"github.com/nmarsollier/cataloggo/graph/tools"
)

func CreateArticle(ctx context.Context, input model.UpdateArticle) (bool, error) {
	_, err := tools.ValidateAdmin(ctx)
	if err != nil {
		return false, err
	}

	env := tools.GqlCtx(ctx)

	_, err = article.CreateArticle(&article.UpdateArticleData{
		Name:        input.Name,
		Description: input.Description,
		Image:       input.Image,
		Price:       float32(input.Price),
		Stock:       input.Stock,
	}, env...)

	if err != nil {
		return false, err
	}

	return true, nil
}

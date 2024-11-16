package resolvers

import (
	"context"

	"github.com/nmarsollier/cataloggo/article"
	"github.com/nmarsollier/cataloggo/graph/model"
	"github.com/nmarsollier/cataloggo/graph/tools"
)

func FindArticleByID(ctx context.Context, articleId string) (*model.Article, error) {
	env := tools.GqlCtx(ctx)

	result, err := article.FindById(articleId, env...)
	if err != nil {
		return nil, err
	}

	return mapArticleDataToModel(result), nil
}

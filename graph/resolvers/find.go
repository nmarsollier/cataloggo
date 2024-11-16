package resolvers

import (
	"context"

	"github.com/nmarsollier/cataloggo/article"
	"github.com/nmarsollier/cataloggo/graph/model"
	"github.com/nmarsollier/cataloggo/graph/tools"
)

func GetArticle(ctx context.Context, articleId string) (*model.Article, error) {
	_, err := tools.ValidateLoggedIn(ctx)
	if err != nil {
		return nil, err
	}

	env := tools.GqlCtx(ctx)

	result, err := article.FindById(articleId, env...)
	if err != nil {
		return nil, err
	}

	return mapArticleDataToModel(result), nil
}

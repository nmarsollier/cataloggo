package resolvers

import (
	"context"

	"github.com/nmarsollier/cataloggo/article"
	"github.com/nmarsollier/cataloggo/graph/model"
	"github.com/nmarsollier/cataloggo/graph/tools"
)

func SearchArticles(ctx context.Context, criteria string) ([]*model.Article, error) {

	_, err := tools.ValidateLoggedIn(ctx)
	if err != nil {
		return nil, err
	}

	env := tools.GqlCtx(ctx)

	articles, err := article.FindByCriteria(criteria, env...)
	if err != nil {
		return nil, err
	}

	result := make([]*model.Article, len(articles))

	for i, a := range articles {
		result[i] = mapArticleDataToModel(a)
	}

	return result, nil
}

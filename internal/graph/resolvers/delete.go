package resolvers

import (
	"context"

	"github.com/nmarsollier/cataloggo/internal/graph/tools"
)

func DeleteArticle(ctx context.Context, articleId string) (bool, error) {
	_, err := tools.ValidateAdmin(ctx)
	if err != nil {
		return false, err
	}

	env := tools.GqlDi(ctx)

	err = env.ArticleService().Disable(articleId)
	if err != nil {
		return false, err
	}

	return true, nil
}

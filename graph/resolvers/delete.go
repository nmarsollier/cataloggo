package resolvers

import (
	"context"

	"github.com/nmarsollier/cataloggo/article"
	"github.com/nmarsollier/cataloggo/graph/tools"
)

func DeleteArticle(ctx context.Context, articleId string) (bool, error) {
	_, err := tools.ValidateAdmin(ctx)
	if err != nil {
		return false, err
	}

	env := tools.GqlCtx(ctx)

	err = article.Disable(articleId, env...)
	if err != nil {
		return false, err
	}

	return true, nil
}

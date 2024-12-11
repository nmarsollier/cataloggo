package article

func UpdateArticle(articleId string, articleData *UpdateArticleData, deps ...interface{}) error {
	err := update(articleId, *articleData, deps...)

	if err != nil {
		return err
	}

	return nil
}

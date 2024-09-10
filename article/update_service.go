package article

func UpdateArticle(articleId string, articleData *UpdateArticleData, ctx ...interface{}) error {
	err := updateDescription(articleId, Description{
		Name:        articleData.Name,
		Description: articleData.Description,
		Image:       articleData.Image,
	}, ctx...)

	if err != nil {
		return err
	}

	err = updateStock(articleId, articleData.Stock, ctx...)

	if err != nil {
		return err
	}

	err = updatePrice(articleId, articleData.Price, ctx...)

	if err != nil {
		return err
	}

	return nil
}

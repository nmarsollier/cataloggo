package article

func UpdateArticle(articleId string, articleData *UpdateArticleData, deps ...interface{}) error {
	err := updateDescription(articleId, Description{
		Name:        articleData.Name,
		Description: articleData.Description,
		Image:       articleData.Image,
	}, deps...)

	if err != nil {
		return err
	}

	err = updateStock(articleId, articleData.Stock, deps...)

	if err != nil {
		return err
	}

	err = updatePrice(articleId, articleData.Price, deps...)

	if err != nil {
		return err
	}

	return nil
}

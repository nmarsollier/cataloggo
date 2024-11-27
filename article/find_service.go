package article

func FindById(id string, deps ...interface{}) (*ArticleData, error) {

	article, err := findById(id, deps...)
	if err != nil {
		return nil, err
	}

	return toArticleData(article), nil
}

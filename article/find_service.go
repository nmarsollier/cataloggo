package article

func FindById(id string, ctx ...interface{}) (*ArticleData, error) {

	article, err := findById(id, ctx...)
	if err != nil {
		return nil, err
	}

	return toArticleData(article), nil
}

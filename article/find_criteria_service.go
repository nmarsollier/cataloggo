package article

func FindByCriteria(criteria string, deps ...interface{}) ([]*ArticleData, error) {
	articles, err := findByCriteria(criteria, deps...)
	if err != nil {
		return nil, err
	}

	result := []*ArticleData{}
	for _, a := range articles {
		result = append(result, toArticleData(a))
	}

	return result, nil
}

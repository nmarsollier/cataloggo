package article

func FindByCriteria(criteria string, ctx ...interface{}) ([]*ArticleData, error) {
	articles, err := findByCriteria(criteria, ctx...)
	if err != nil {
		return nil, err
	}

	result := []*ArticleData{}
	for _, a := range articles {
		result = append(result, toArticleData(a))
	}

	return result, nil
}

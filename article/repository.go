package article

import (
	"github.com/nmarsollier/cataloggo/tools/errs"
	"github.com/nmarsollier/cataloggo/tools/log"
)

var ErrID = errs.NewValidation().Add("id", "Invalid")

func findByCriteria(criteria string, deps ...interface{}) (result []*Article, err error) {
	connection, err := GetArticleDao(deps...)
	if err != nil {
		log.Get(deps...).Error(err)
		return
	}

	result, err = connection.FindByCriteria(criteria)
	if err != nil {
		log.Get(deps...).Error(err)
	}

	return
}

type DBCriteriaFilter struct {
	Or []map[string]DBCriteriaElement `bson:"$or"`
}

type DBCriteriaElement struct {
	RegEx   string `bson:"$regex"`
	Options string `bson:"$options"`
}

func findById(articleId string, deps ...interface{}) (result *Article, err error) {
	connection, err := GetArticleDao(deps...)
	if err != nil {
		log.Get(deps...).Error(err)
		return
	}

	result, err = connection.FindById(articleId)
	if err != nil {
		log.Get(deps...).Error(err)
	}

	return
}

func insert(article *Article, deps ...interface{}) (err error) {
	if err = article.validateSchema(); err != nil {
		log.Get(deps...).Error(err)

		return
	}

	connection, err := GetArticleDao(deps...)
	if err != nil {
		log.Get(deps...).Error(err)

		return
	}

	if err = connection.Insert(article); err != nil {
		log.Get(deps...).Error(err)
		return
	}

	return
}

// disable Deshabilita el articulo para que no se pueda usar mas
func Disable(articleId string, deps ...interface{}) (err error) {
	connection, err := GetArticleDao(deps...)
	if err != nil {
		log.Get(deps...).Error(err)

		return
	}

	err = connection.Disable(articleId)

	return
}

// Actualiza la descripción de un articulo.
func updateDescription(articleId string, description Description, deps ...interface{}) (err error) {
	collection, err := GetArticleDao(deps...)
	if err != nil {
		log.Get(deps...).Error(err)

		return
	}

	err = collection.UpdateDescription(articleId, description)

	return
}

// Actualiza el precio de un articulo.
func updatePrice(articleId string, price float32, deps ...interface{}) (err error) {

	collection, err := GetArticleDao(deps...)
	if err != nil {
		log.Get(deps...).Error(err)
		return
	}

	err = collection.UpdatePrice(articleId, price)

	return
}

// Actualiza el stock de un articulo.
func updateStock(articleId string, stock int, deps ...interface{}) (err error) {
	collection, err := GetArticleDao(deps...)
	if err != nil {
		log.Get(deps...).Error(err)

		return
	}

	err = collection.UpdateStock(articleId, stock)

	if err != nil {
		log.Get(deps...).Error(err)
	}

	return
}

func DecrementStock(articleId string, amount int, deps ...interface{}) (err error) {
	collection, err := GetArticleDao(deps...)
	if err != nil {
		log.Get(deps...).Error(err)

		return
	}

	err = collection.DecrementStock(articleId, amount)

	if err != nil {
		log.Get(deps...).Error(err)
	}

	return
}

package article

import (
	"context"

	"github.com/golang/glog"
	"github.com/nmarsollier/cataloggo/tools/apperr"
	"github.com/nmarsollier/cataloggo/tools/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Define mongo Collection
var collection *mongo.Collection

func dbCollection() (*mongo.Collection, error) {
	if collection != nil {
		return collection, nil
	}

	database, err := db.Get()
	if err != nil {
		glog.Error(err)

		return nil, err
	}

	col := database.Collection("articles")

	collection = col
	return collection, nil
}

func findByCriteria(criteria string) ([]*Article, error) {
	var collection, err = dbCollection()
	if err != nil {
		glog.Error(err)

		return nil, err
	}

	filter := bson.M{
		"$or": []bson.M{
			{"description.name": bson.M{"$regex": criteria, "$options": "i"}},
			{"description.description": bson.M{"$regex": criteria, "$options": "i"}},
		},
	}

	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		glog.Error(err)

		return nil, err
	}
	defer cur.Close(context.Background())

	result := []*Article{}
	for cur.Next(context.Background()) {
		article := &Article{}
		if err := cur.Decode(article); err != nil {
			glog.Error(err)

			return nil, err
		}
		result = append(result, article)
	}

	return result, nil
}

func findById(articleId string) (*Article, error) {
	var collection, err = dbCollection()
	if err != nil {
		glog.Error(err)

		return nil, err
	}

	_id, err := primitive.ObjectIDFromHex(articleId)
	if err != nil {
		glog.Error(err)

		return nil, apperr.ErrID
	}

	article := &Article{}
	filter := bson.M{"_id": _id}

	if err = collection.FindOne(context.Background(), filter).Decode(article); err != nil {
		glog.Error(err)

		return nil, err
	}

	return article, nil
}

func insert(article *Article) (*Article, error) {
	if err := article.ValidateSchema(); err != nil {
		glog.Error(err)

		return nil, err
	}

	var collection, err = dbCollection()
	if err != nil {
		glog.Error(err)

		return nil, err
	}

	if _, err = collection.InsertOne(context.Background(), article); err != nil {
		glog.Error(err)

		return nil, err
	}

	return article, nil
}

// disable Deshabilita el articulo para que no se pueda usar mas
func Disable(articleId string) error {
	var collection, err = dbCollection()
	if err != nil {
		glog.Error(err)

		return err
	}

	_id, err := primitive.ObjectIDFromHex(articleId)
	if err != nil {
		glog.Error(err)

		return apperr.ErrID
	}

	_, err = collection.UpdateOne(context.Background(),
		bson.M{"_id": _id},
		bson.M{"$set": bson.M{
			"enabled": false,
		}},
	)

	return err
}

// Actualiza la descripción de un articulo.
func updateDescription(articleId string, description Description) error {
	var collection, err = dbCollection()
	if err != nil {
		glog.Error(err)

		return err
	}

	_id, err := primitive.ObjectIDFromHex(articleId)
	if err != nil {
		glog.Error(err)

		return apperr.ErrID
	}
	_, err = collection.UpdateOne(context.Background(),
		bson.M{"_id": _id},
		bson.M{"$set": bson.M{
			"description": description,
		}},
	)

	return err
}

// Actualiza el precio de un articulo.
func updatePrice(articleId string, price float32) error {
	var collection, err = dbCollection()
	if err != nil {
		glog.Error(err)

		return err
	}

	_id, err := primitive.ObjectIDFromHex(articleId)
	if err != nil {
		glog.Error(err)

		return apperr.ErrID
	}
	_, err = collection.UpdateOne(context.Background(),
		bson.M{"_id": _id},
		bson.M{"$set": bson.M{
			"price": price,
		}},
	)

	return err
}

// Actualiza el stock de un articulo.
func updateStock(articleId string, stock int) error {
	var collection, err = dbCollection()
	if err != nil {
		glog.Error(err)

		return err
	}

	_id, err := primitive.ObjectIDFromHex(articleId)
	if err != nil {
		glog.Error(err)

		return apperr.ErrID
	}
	_, err = collection.UpdateOne(context.Background(),
		bson.M{"_id": _id},
		bson.M{"$set": bson.M{
			"stock": stock,
		}},
	)

	return err
}

func DecreaseStock(articleId primitive.ObjectID, stock int) error {
	var collection, err = dbCollection()
	if err != nil {
		glog.Error(err)

		return err
	}

	_, err = collection.UpdateOne(context.Background(),
		bson.M{"_id": articleId},
		bson.M{
			"$inc": bson.M{
				"stock": -stock,
			}},
	)

	return err
}

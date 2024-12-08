package article

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/nmarsollier/cataloggo/tools/db"
	"github.com/nmarsollier/cataloggo/tools/errs"
	"github.com/nmarsollier/cataloggo/tools/log"
)

var tableName = "articles"

var ErrID = errs.NewValidation().Add("id", "Invalid")

func findByCriteria(criteria string, deps ...interface{}) (result []*Article, err error) {
	filterExpression := "contains(description.#name, :criteria) OR contains(description.#description, :criteria)"
	expressionAttributeNames := map[string]string{
		"#name":        "name",
		"#description": "description",
	}
	expressionAttributeValues := map[string]types.AttributeValue{
		":criteria": &types.AttributeValueMemberS{
			Value: criteria,
		},
	}

	input := &dynamodb.ScanInput{
		TableName:                 aws.String(tableName),
		FilterExpression:          aws.String(filterExpression),
		ExpressionAttributeNames:  expressionAttributeNames,
		ExpressionAttributeValues: expressionAttributeValues,
	}

	output, err := db.Get(deps...).Scan(context.TODO(), input)
	if err != nil {
		log.Get(deps...).Error(err)

		return
	}

	err = attributevalue.UnmarshalListOfMaps(output.Items, &result)
	if err != nil {
		log.Get(deps...).Error(err)
	}

	return
}

func findById(articleId string, deps ...interface{}) (result *Article, err error) {
	response, err := db.Get(deps...).GetItem(context.TODO(), &dynamodb.GetItemInput{
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{
				Value: articleId,
			}},
		TableName: &tableName,
	})

	if err != nil || response == nil || response.Item == nil {
		log.Get(deps...).Error(err)

		return nil, errs.NotFound
	}

	err = attributevalue.UnmarshalMap(response.Item, &result)
	if err != nil {
		log.Get(deps...).Error(err)

		return
	}

	return
}

func insert(article *Article, deps ...interface{}) (err error) {
	if err = article.validateSchema(); err != nil {
		log.Get(deps...).Error(err)

		return
	}

	articleToInsert, err := attributevalue.MarshalMap(article)
	if err != nil {
		log.Get(deps...).Error(err)

		return
	}

	_, err = db.Get(deps...).PutItem(
		context.TODO(),
		&dynamodb.PutItemInput{
			TableName: &tableName,
			Item:      articleToInsert,
		},
	)
	return
}

// disable Deshabilita el articulo para que no se pueda usar mas
func Disable(articleId string, deps ...interface{}) (err error) {
	key, err := attributevalue.MarshalMap(map[string]interface{}{
		"id": articleId,
	})
	if err != nil {
		log.Get(deps...).Error(err)

		return
	}

	update, err := attributevalue.MarshalMap(map[string]interface{}{
		":enabled": false,
	})
	if err != nil {
		log.Get(deps...).Error(err)

		return
	}

	_, err = db.Get(deps...).UpdateItem(
		context.TODO(),
		&dynamodb.UpdateItemInput{
			TableName:                 &tableName,
			Key:                       key,
			UpdateExpression:          aws.String("SET enabled = :enabled"),
			ExpressionAttributeValues: update,
		},
	)

	if err != nil {
		log.Get(deps...).Error(err)
	}

	return
}

// Actualiza la descripción de un articulo.
func updateDescription(articleId string, description Description, deps ...interface{}) (err error) {
	key, err := attributevalue.MarshalMap(map[string]interface{}{
		"id": articleId,
	})
	if err != nil {
		log.Get(deps...).Error(err)

		return
	}

	update, err := attributevalue.MarshalMap(map[string]interface{}{
		":description": description,
	})
	if err != nil {
		log.Get(deps...).Error(err)

		return
	}

	_, err = db.Get(deps...).UpdateItem(
		context.TODO(),
		&dynamodb.UpdateItemInput{
			TableName:                 &tableName,
			Key:                       key,
			UpdateExpression:          aws.String("SET description = :description"),
			ExpressionAttributeValues: update,
		},
	)

	if err != nil {
		log.Get(deps...).Error(err)
	}
	return
}

// Actualiza el precio de un articulo.
func updatePrice(articleId string, price float32, deps ...interface{}) (err error) {
	key, err := attributevalue.MarshalMap(map[string]interface{}{
		"id": articleId,
	})
	if err != nil {
		log.Get(deps...).Error(err)

		return
	}

	update, err := attributevalue.MarshalMap(map[string]interface{}{
		":price": price,
	})
	if err != nil {
		log.Get(deps...).Error(err)

		return
	}

	_, err = db.Get(deps...).UpdateItem(
		context.TODO(),
		&dynamodb.UpdateItemInput{
			TableName:                 &tableName,
			Key:                       key,
			UpdateExpression:          aws.String("SET price = :price"),
			ExpressionAttributeValues: update,
		},
	)
	if err != nil {
		log.Get(deps...).Error(err)
	}
	return
}

// Actualiza el stock de un articulo.
func updateStock(articleId string, stock int, deps ...interface{}) (err error) {
	key, err := attributevalue.MarshalMap(map[string]interface{}{
		"id": articleId,
	})
	if err != nil {
		log.Get(deps...).Error(err)

		return err
	}

	update, err := attributevalue.MarshalMap(map[string]interface{}{
		":stock": stock,
	})
	if err != nil {
		log.Get(deps...).Error(err)

		return err
	}

	_, err = db.Get(deps...).UpdateItem(
		context.TODO(),
		&dynamodb.UpdateItemInput{
			TableName:                 &tableName,
			Key:                       key,
			UpdateExpression:          aws.String("SET stock = :stock"),
			ExpressionAttributeValues: update,
		},
	)
	if err != nil {
		log.Get(deps...).Error(err)
	}
	return nil
}

func DecrementStock(articleId string, amount int, deps ...interface{}) (err error) {
	key, err := attributevalue.MarshalMap(map[string]interface{}{
		"id": articleId,
	})
	if err != nil {
		log.Get(deps...).Error(err)

		return err
	}

	update, err := attributevalue.MarshalMap(map[string]interface{}{
		":decrement": amount,
	})
	if err != nil {
		log.Get(deps...).Error(err)

		return err
	}

	_, err = db.Get(deps...).UpdateItem(
		context.TODO(),
		&dynamodb.UpdateItemInput{
			TableName:                 &tableName,
			Key:                       key,
			UpdateExpression:          aws.String("SET stock = stock - :decrement"),
			ExpressionAttributeValues: update,
		},
	)
	if err != nil {
		log.Get(deps...).Error(err)

	}
	return
}

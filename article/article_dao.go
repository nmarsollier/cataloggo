package article

import (
	"context"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/nmarsollier/cataloggo/tools/env"
)

var tableName = "articles"

var (
	once     sync.Once
	instance ArticleDao
)

type ArticleDao interface {
	FindByCriteria(criteria string) (result []*Article, err error)

	FindById(key string) (*Article, error)

	Insert(article *Article) error

	Disable(articleId string) error

	UpdateDescription(articleId string, description Description) error

	UpdatePrice(articleId string, price float32) error

	UpdateStock(articleId string, stock int) error

	DecrementStock(articleId string, amount int) error
}

func GetArticleDao(deps ...interface{}) (ArticleDao, error) {
	for _, o := range deps {
		if client, ok := o.(ArticleDao); ok {
			return client, nil
		}
	}

	var conn_err error
	once.Do(func() {
		customCreds := aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(
			env.Get().AwsAccessKeyId,
			env.Get().AwsSecret,
			"",
		))

		cfg, err := config.LoadDefaultConfig(context.TODO(),
			config.WithRegion(env.Get().AwsRegion),
			config.WithCredentialsProvider(customCreds),
		)
		if err != nil {
			conn_err = err
			return
		}

		instance = &articleDao{
			client: dynamodb.NewFromConfig(cfg),
		}
	})

	return instance, conn_err
}

type articleDao struct {
	client *dynamodb.Client
}

func (r *articleDao) FindByCriteria(criteria string) (result []*Article, err error) {
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

	output, err := r.client.Scan(context.TODO(), input)
	if err != nil {
		return nil, err
	}

	var articles []*Article
	err = attributevalue.UnmarshalListOfMaps(output.Items, &articles)
	if err != nil {
		return nil, err
	}

	return articles, nil
}

func (r *articleDao) FindById(key string) (*Article, error) {
	user := Article{ID: key}
	userId, err := attributevalue.Marshal(user.ID)
	if err != nil {
		return nil, err
	}

	response, err := r.client.GetItem(context.TODO(), &dynamodb.GetItemInput{
		Key: map[string]types.AttributeValue{"id": userId}, TableName: &tableName,
	})

	if err != nil || response == nil || response.Item == nil {
		return nil, err
	}

	err = attributevalue.UnmarshalMap(response.Item, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *articleDao) Insert(article *Article) error {
	articleToInsert, err := attributevalue.MarshalMap(article)
	if err != nil {
		return err
	}

	_, err = r.client.PutItem(
		context.TODO(),
		&dynamodb.PutItemInput{
			TableName: &tableName,
			Item:      articleToInsert,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *articleDao) Disable(articleId string) error {
	key, err := attributevalue.MarshalMap(map[string]interface{}{
		"id": articleId,
	})
	if err != nil {
		return err
	}

	update, err := attributevalue.MarshalMap(map[string]interface{}{
		":enabled": false,
	})
	if err != nil {
		return err
	}

	_, err = r.client.UpdateItem(
		context.TODO(),
		&dynamodb.UpdateItemInput{
			TableName:                 &tableName,
			Key:                       key,
			UpdateExpression:          aws.String("SET enabled = :enabled"),
			ExpressionAttributeValues: update,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *articleDao) UpdateDescription(articleId string, description Description) error {
	key, err := attributevalue.MarshalMap(map[string]interface{}{
		"id": articleId,
	})
	if err != nil {
		return err
	}

	update, err := attributevalue.MarshalMap(map[string]interface{}{
		":description": description,
	})
	if err != nil {
		return err
	}

	_, err = r.client.UpdateItem(
		context.TODO(),
		&dynamodb.UpdateItemInput{
			TableName:                 &tableName,
			Key:                       key,
			UpdateExpression:          aws.String("SET description = :description"),
			ExpressionAttributeValues: update,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *articleDao) UpdatePrice(articleId string, price float32) error {
	key, err := attributevalue.MarshalMap(map[string]interface{}{
		"id": articleId,
	})
	if err != nil {
		return err
	}

	update, err := attributevalue.MarshalMap(map[string]interface{}{
		":price": price,
	})
	if err != nil {
		return err
	}

	_, err = r.client.UpdateItem(
		context.TODO(),
		&dynamodb.UpdateItemInput{
			TableName:                 &tableName,
			Key:                       key,
			UpdateExpression:          aws.String("SET price = :price"),
			ExpressionAttributeValues: update,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *articleDao) UpdateStock(articleId string, stock int) error {
	key, err := attributevalue.MarshalMap(map[string]interface{}{
		"id": articleId,
	})
	if err != nil {
		return err
	}

	update, err := attributevalue.MarshalMap(map[string]interface{}{
		":stock": stock,
	})
	if err != nil {
		return err
	}

	_, err = r.client.UpdateItem(
		context.TODO(),
		&dynamodb.UpdateItemInput{
			TableName:                 &tableName,
			Key:                       key,
			UpdateExpression:          aws.String("SET stock = :stock"),
			ExpressionAttributeValues: update,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *articleDao) DecrementStock(articleId string, amount int) error {
	key, err := attributevalue.MarshalMap(map[string]interface{}{
		"id": articleId,
	})
	if err != nil {
		return err
	}

	update, err := attributevalue.MarshalMap(map[string]interface{}{
		":decrement": amount,
	})
	if err != nil {
		return err
	}

	_, err = r.client.UpdateItem(
		context.TODO(),
		&dynamodb.UpdateItemInput{
			TableName:                 &tableName,
			Key:                       key,
			UpdateExpression:          aws.String("SET stock = stock - :decrement"),
			ExpressionAttributeValues: update,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

package article

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type PaymentMethod string

// Estuctura basica de del evento
type Article struct {
	ID          string      `dynamodbav:"id" json:"id"`
	Description Description `dynamodbav:"description"  json:"description" validate:"required"`
	Price       float32     `dynamodbav:"price"  json:"price"`
	Stock       int         `dynamodbav:"stock"  json:"stock"`
	Created     time.Time   `dynamodbav:"created" json:"created"`
	Updated     time.Time   `dynamodbav:"updated" json:"updated"`
	Enabled     bool        `dynamodbav:"enabled" json:"enabled"`
}

type Description struct {
	Name        string `dynamodbav:"name"  json:"name" validate:"required,min=1,max=100"`
	Description string `dynamodbav:"description"  json:"description" validate:"required,min=1,max=256"`
	Image       string `dynamodbav:"image"  json:"image" validate:"max=100"`
}

// validateSchema valida la estructura para ser insertada en la db
func (e *Article) validateSchema() error {
	return validator.New().Struct(e)
}

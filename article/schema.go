package article

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type PaymentMethod string

// Estuctura basica de del evento
type Article struct {
	ID          string    `json:"id"`
	Name        string    `json:"name" validate:"required,min=1,max=100"`
	Description string    `json:"description" validate:"required,min=1,max=256"`
	Image       string    `json:"image" validate:"max=100"`
	Price       float32   `json:"price"`
	Stock       int       `json:"stock"`
	Created     time.Time `json:"created"`
	Updated     time.Time `json:"updated"`
	Enabled     bool      `json:"enabled"`
}

// validateSchema valida la estructura para ser insertada en la db
func (e *Article) validateSchema() error {
	return validator.New().Struct(e)
}

// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Article struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
}

func (Article) IsEntity() {}

type CreateArticle struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
}

type Mutation struct {
}

type Query struct {
}

type UpdateArticle struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
}

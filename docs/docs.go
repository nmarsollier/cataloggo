// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Nestor Marsollier",
            "email": "nmarsollier@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/rabbit/article_exist": {
            "get": {
                "description": "Otros microservicios nos solicitan validar articulos en el catalogo.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Rabbit"
                ],
                "summary": "Mensage Rabbit article_exist/article_exist",
                "parameters": [
                    {
                        "description": "Message para article_exist",
                        "name": "article_exist",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/rschema.ConsumeArticleExist"
                        }
                    }
                ],
                "responses": {}
            },
            "put": {
                "description": "Emite respuestas de article_exist",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Rabbit"
                ],
                "summary": "Mensage Rabbit article_exist",
                "parameters": [
                    {
                        "description": "Estructura general del mensage",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/rschema.SendArticleExist"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/rabbit/logout": {
            "get": {
                "description": "Escucha de mensajes logout desde auth.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Rabbit"
                ],
                "summary": "Mensage Rabbit logout",
                "parameters": [
                    {
                        "description": "Estructura general del mensage",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/consume.logoutMessage"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/rabbit/order_placed": {
            "get": {
                "description": "Cuando se recibe el mensage order_placed damos de baja al stock para reservar los articulos. Queda pendiente enviar mensaje confirmando la operacion al MS de Orders.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Rabbit"
                ],
                "summary": "Mensage Rabbit order_placed/order_placed",
                "parameters": [
                    {
                        "description": "Message order_placed",
                        "name": "order_placed",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.ConsumeOrderPlacedMessage"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/v1/articles": {
            "post": {
                "description": "Crear Artículo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Catalogo"
                ],
                "summary": "Crear Artículo",
                "parameters": [
                    {
                        "type": "string",
                        "description": "bearer {token}",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Informacion del articulo",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/article.NewArticleData"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Articulo",
                        "schema": {
                            "$ref": "#/definitions/article.ArticleData"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errs.ValidationErr"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/engine.ErrorData"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/engine.ErrorData"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/engine.ErrorData"
                        }
                    }
                }
            }
        },
        "/v1/articles/:articleId": {
            "get": {
                "description": "Obtener un articulo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Catalogo"
                ],
                "summary": "Obtener un articulo",
                "parameters": [
                    {
                        "type": "string",
                        "description": "bearer {token}",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "ID de articlo",
                        "name": "articleId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Articulo",
                        "schema": {
                            "$ref": "#/definitions/article.ArticleData"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errs.ValidationErr"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/engine.ErrorData"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/engine.ErrorData"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/engine.ErrorData"
                        }
                    }
                }
            },
            "post": {
                "description": "Actualizar Artículo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Catalogo"
                ],
                "summary": "Actualizar Artículo",
                "parameters": [
                    {
                        "type": "string",
                        "description": "bearer {token}",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Informacion del articulo",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/article.NewArticleData"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Articulo",
                        "schema": {
                            "$ref": "#/definitions/article.ArticleData"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errs.ValidationErr"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/engine.ErrorData"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/engine.ErrorData"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/engine.ErrorData"
                        }
                    }
                }
            },
            "delete": {
                "description": "Eliminar Artículo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Catalogo"
                ],
                "summary": "Eliminar Artículo",
                "parameters": [
                    {
                        "type": "string",
                        "description": "bearer {token}",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "ID de articlo",
                        "name": "articleId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errs.ValidationErr"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/engine.ErrorData"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/engine.ErrorData"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/engine.ErrorData"
                        }
                    }
                }
            }
        },
        "/v1/articles/search/:criteria": {
            "get": {
                "description": "Obtener un articulo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Catalogo"
                ],
                "summary": "Obtener un articulo",
                "parameters": [
                    {
                        "type": "string",
                        "description": "bearer {token}",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "ID de articlo",
                        "name": "articleId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Articulos",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/article.ArticleData"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errs.ValidationErr"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/engine.ErrorData"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/engine.ErrorData"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/engine.ErrorData"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "article.ArticleData": {
            "type": "object",
            "properties": {
                "_id": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "enabled": {
                    "type": "boolean"
                },
                "image": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                },
                "stock": {
                    "type": "integer"
                }
            }
        },
        "article.NewArticleData": {
            "type": "object",
            "required": [
                "description",
                "name",
                "price",
                "stock"
            ],
            "properties": {
                "description": {
                    "type": "string",
                    "maxLength": 256,
                    "minLength": 1
                },
                "image": {
                    "type": "string",
                    "maxLength": 100
                },
                "name": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 1
                },
                "price": {
                    "type": "number",
                    "minimum": 1
                },
                "stock": {
                    "type": "integer",
                    "minimum": 1
                }
            }
        },
        "consume.logoutMessage": {
            "type": "object",
            "properties": {
                "correlation_id": {
                    "type": "string",
                    "example": "123123"
                },
                "message": {
                    "type": "string",
                    "example": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0b2tlbklEIjoiNjZiNjBlYzhlMGYzYzY4OTUzMzJlOWNmIiwidXNlcklEIjoiNjZhZmQ3ZWU4YTBhYjRjZjQ0YTQ3NDcyIn0.who7upBctOpmlVmTvOgH1qFKOHKXmuQCkEjMV3qeySg"
                }
            }
        },
        "engine.ErrorData": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "errs.ValidationErr": {
            "type": "object",
            "properties": {
                "messages": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/errs.errField"
                    }
                }
            }
        },
        "errs.errField": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "path": {
                    "type": "string"
                }
            }
        },
        "rschema.ArticleExistMessage": {
            "type": "object",
            "properties": {
                "articleId": {
                    "type": "string",
                    "example": "ArticleId"
                },
                "price": {
                    "type": "number"
                },
                "referenceId": {
                    "type": "string",
                    "example": "Remote Reference Id"
                },
                "stock": {
                    "type": "integer"
                },
                "valid": {
                    "type": "boolean"
                }
            }
        },
        "rschema.ConsumeArticleExist": {
            "type": "object",
            "properties": {
                "correlation_id": {
                    "type": "string",
                    "example": "123123"
                },
                "exchange": {
                    "type": "string",
                    "example": "Remote Exchange to Reply"
                },
                "message": {
                    "$ref": "#/definitions/rschema.ConsumeArticleExistMessage"
                },
                "routing_key": {
                    "type": "string",
                    "example": "Remote RoutingKey to Reply"
                }
            }
        },
        "rschema.ConsumeArticleExistMessage": {
            "type": "object",
            "properties": {
                "articleId": {
                    "type": "string",
                    "example": "ArticleId"
                },
                "referenceId": {
                    "type": "string",
                    "example": "Remote Reference Object Id"
                }
            }
        },
        "rschema.SendArticleExist": {
            "type": "object",
            "properties": {
                "correlation_id": {
                    "type": "string",
                    "example": "123123"
                },
                "message": {
                    "$ref": "#/definitions/rschema.ArticleExistMessage"
                }
            }
        },
        "service.ConsumeOrderPlacedArticle": {
            "type": "object",
            "properties": {
                "articleId": {
                    "type": "string"
                },
                "quantity": {
                    "type": "integer"
                }
            }
        },
        "service.ConsumeOrderPlacedMessage": {
            "type": "object",
            "properties": {
                "articles": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/service.ConsumeOrderPlacedArticle"
                    }
                },
                "cartId": {
                    "type": "string"
                },
                "orderId": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3002",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "CatalogGo",
	Description:      "Microservicio de Catalogo.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/cataloggo/article"
	"github.com/nmarsollier/cataloggo/rest/engine"
	"github.com/nmarsollier/cataloggo/rest/middlewares"
)

// Obtener un articulo
//
//	@Summary		Obtener un articulo
//	@Description	Obtener un articulo
//	@Tags			Catalogo
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string					true	"bearer {token}"
//	@Param			articleId		path		string					true	"ID de articlo"
//	@Success		200				{object}	article.ArticleData		"Articulo"
//	@Failure		400				{object}	apperr.ErrValidation	"Bad Request"
//	@Failure		401				{object}	apperr.ErrCustom		"Unauthorized"
//	@Failure		404				{object}	apperr.ErrCustom		"Not Found"
//	@Failure		500				{object}	apperr.ErrCustom		"Internal Server Error"
//
//	@Router			/v1/articles/:articleId [get]
func init() {
	engine.Router().GET(
		"/v1/articles/:articleId",
		middlewares.ValidateAuthentication,
		getArticle,
	)
}

func getArticle(c *gin.Context) {
	articleId := c.Param("articleId")

	result, err := article.FindById(articleId)
	if err != nil {
		middlewares.AbortWithError(c, err)
		return
	}

	c.JSON(200, result)
}

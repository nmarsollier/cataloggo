package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/cataloggo/article"
	"github.com/nmarsollier/cataloggo/rest/engine"
	"github.com/nmarsollier/cataloggo/rest/middlewares"
)

// Eliminar Artículo
//
//	@Summary		Eliminar Artículo
//	@Description	Eliminar Artículo
//	@Tags			Catalogo
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header	string	true	"bearer {token}"
//	@Param			articleId		path	string	true	"ID de articlo"
//	@Success		200				"No Content"
//	@Failure		400				{object}	apperr.ErrValidation	"Bad Request"
//	@Failure		401				{object}	apperr.ErrCustom		"Unauthorized"
//	@Failure		404				{object}	apperr.ErrCustom		"Not Found"
//	@Failure		500				{object}	apperr.ErrCustom		"Internal Server Error"
//
//	@Router			/v1/articles/:articleId [delete]
func init() {
	engine.Router().DELETE(
		"/v1/articles/:articleId",
		middlewares.ValidateAuthentication,
		deleteArticle,
	)
}

func deleteArticle(c *gin.Context) {
	articleId := c.Param("articleId")

	err := article.Disable(articleId)
	if err != nil {
		middlewares.AbortWithError(c, err)
		return
	}

	c.JSON(200, "")
}

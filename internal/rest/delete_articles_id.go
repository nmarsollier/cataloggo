package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/cataloggo/internal/rest/engine"
)

//	@Summary		Eliminar Artículo
//	@Description	Eliminar Artículo
//	@Tags			Catalogo
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header	string	true	"Bearer {token}"
//	@Param			articleId		path	string	true	"ID de articlo"
//	@Success		200				"No Content"
//	@Failure		400				{object}	errs.ValidationErr	"Bad Request"
//	@Failure		401				{object}	engine.ErrorData	"Unauthorized"
//	@Failure		404				{object}	engine.ErrorData	"Not Found"
//	@Failure		500				{object}	engine.ErrorData	"Internal Server Error"
//	@Router			/articles/:articleId [delete]
//
// Eliminar Artículo
func init() {
	engine.Router().DELETE(
		"/articles/:articleId",
		engine.ValidateAuthentication,
		deleteArticle,
	)
}

func deleteArticle(c *gin.Context) {
	articleId := c.Param("articleId")
	deps := engine.GinDi(c)

	err := deps.ArticleService().Disable(articleId)
	if err != nil {
		engine.AbortWithError(c, err)
		return
	}

	c.JSON(200, "")
}

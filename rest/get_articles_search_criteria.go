package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/cataloggo/article"
	"github.com/nmarsollier/cataloggo/rest/engine"
)

//	@Summary		Obtener un articulo
//	@Description	Obtener un articulo
//	@Tags			Catalogo
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string				true	"Bearer {token}"
//	@Param			articleId		path		string				true	"ID de articlo"
//	@Success		200				{array}		article.ArticleData	"Articulos"
//	@Failure		400				{object}	errs.ValidationErr	"Bad Request"
//	@Failure		401				{object}	engine.ErrorData	"Unauthorized"
//	@Failure		404				{object}	engine.ErrorData	"Not Found"
//	@Failure		500				{object}	engine.ErrorData	"Internal Server Error"
//	@Router			/v1/articles/search/:criteria [get]
//
// Obtener un articulo
func init() {
	engine.Router().GET(
		"/v1/articles/search/:criteria",
		engine.ValidateAuthentication,
		getArticleSearch,
	)
}

func getArticleSearch(c *gin.Context) {
	criteria := c.Param("criteria")

	deps := engine.GinDeps(c)

	result, err := article.FindByCriteria(criteria, deps...)
	if err != nil {
		engine.AbortWithError(c, err)
		return
	}

	c.JSON(200, result)
}

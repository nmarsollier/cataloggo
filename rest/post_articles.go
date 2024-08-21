package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/cataloggo/article"
	"github.com/nmarsollier/cataloggo/rest/engine"
)

//	@Summary		Crear Artículo
//	@Description	Crear Artículo
//	@Tags			Catalogo
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string					true	"bearer {token}"
//	@Param			body			body		article.NewArticleData	true	"Informacion del articulo"
//	@Success		200				{object}	article.ArticleData		"Articulo"
//	@Failure		400				{object}	errs.ValidationErr		"Bad Request"
//	@Failure		401				{object}	engine.ErrorData		"Unauthorized"
//	@Failure		404				{object}	engine.ErrorData		"Not Found"
//	@Failure		500				{object}	engine.ErrorData		"Internal Server Error"
//	@Router			/v1/articles [post]
//
// Crear Artículo
func init() {
	engine.Router().POST(
		"/v1/articles",
		engine.ValidateAuthentication,
		saveArticle,
	)
}

func saveArticle(c *gin.Context) {
	body := article.NewArticleData{}
	if err := c.ShouldBindJSON(&body); err != nil {
		engine.AbortWithError(c, err)
		return
	}

	ctx := engine.GinCtx(c)
	article, err := article.CreateArticle(&body, ctx...)
	if err != nil {
		engine.AbortWithError(c, err)
		return
	}

	c.JSON(200, article)
}

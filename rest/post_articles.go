package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/cataloggo/article"
	"github.com/nmarsollier/cataloggo/rest/engine"
	"github.com/nmarsollier/cataloggo/rest/middlewares"
)

//	Crear Artículo
//
// @Summary		Crear Artículo
// @Description	Crear Artículo
// @Tags			Catalogo
// @Accept			json
// @Produce		json
// @Param			Authorization	header		string					true	"bearer {token}"
// @Param			body			body		article.NewArticleData	true	"Informacion del articulo"
// @Success		200				{object}	article.ArticleData		"Articulo"
// @Failure		400				{object}	apperr.ErrValidation	"Bad Request"
// @Failure		401				{object}	apperr.ErrCustom		"Unauthorized"
// @Failure		404				{object}	apperr.ErrCustom		"Not Found"
// @Failure		500				{object}	apperr.ErrCustom		"Internal Server Error"
//
// @Router			/v1/articles [post]
func init() {
	engine.Router().POST(
		"/v1/articles",
		middlewares.ValidateAuthentication,
		saveArticle,
	)
}

func saveArticle(c *gin.Context) {
	body := article.NewArticleData{}
	if err := c.ShouldBindJSON(&body); err != nil {
		middlewares.AbortWithError(c, err)
		return
	}

	article, err := article.CreateArticle(&body)
	if err != nil {
		middlewares.AbortWithError(c, err)
		return
	}

	c.JSON(200, article)
}

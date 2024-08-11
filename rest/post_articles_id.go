package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/cataloggo/article"
	"github.com/nmarsollier/cataloggo/rest/engine"
	"github.com/nmarsollier/cataloggo/rest/middlewares"
)

//	Actualizar Artículo
//
// @Summary		Actualizar Artículo
// @Description	Actualizar Artículo
// @Tags			Catalogo
// @Accept			json
// @Produce		json
// @Param			Authorization	header		string					true	"bearer {token}"
// @Param			body			body		article.NewArticleData	true	"Informacion del articulo"
// @Success		200				{object}	article.ArticleData		"Articulo"
// @Failure		400				{object}	errors.ErrValidation	"Bad Request"
// @Failure		401				{object}	errors.ErrCustom		"Unauthorized"
// @Failure		404				{object}	errors.ErrCustom		"Not Found"
// @Failure		500				{object}	errors.ErrCustom		"Internal Server Error"
//
// @Router			/v1/articles/:articleId [post]
func init() {
	engine.Router().POST(
		"/v1/articles/:articleId",
		middlewares.ValidateAuthentication,
		updateArticle,
	)
}

func updateArticle(c *gin.Context) {
	body := article.NewArticleData{}
	if err := c.ShouldBindJSON(&body); err != nil {
		middlewares.AbortWithError(c, err)
		return
	}
	articleId := c.Param("articleId")

	err := article.UpdateArticle(articleId, &body)
	if err != nil {
		middlewares.AbortWithError(c, err)
		return
	}

	c.JSON(200, "")
}

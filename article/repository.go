package article

import (
	"context"
	"fmt"

	"github.com/jackc/pgx"
	"github.com/nmarsollier/cataloggo/tools/db"
	"github.com/nmarsollier/cataloggo/tools/errs"
	"github.com/nmarsollier/cataloggo/tools/log"
)

var ErrID = errs.NewValidation().Add("id", "Invalid")

func findByCriteria(criteria string, deps ...interface{}) (result []*Article, err error) {
	conn, err := db.GetPostgresClient(deps...)
	if err != nil {
		log.Get(deps...).Error(err)
		return nil, err
	}

	query := `
        SELECT id, name, description, image, price, stock, created, updated, enabled
        FROM Articles
        WHERE name ILIKE $1 OR description ILIKE $1
    `

	rows, err := conn.Query(context.Background(), query, fmt.Sprintf("%%%s%%", criteria))
	if err != nil {
		log.Get(deps...).Error(err)
		return nil, err
	}
	defer rows.Close()

	var articles []*Article
	for rows.Next() {
		var article Article
		err := rows.Scan(
			&article.ID,
			&article.Name,
			&article.Description,
			&article.Image,
			&article.Price,
			&article.Stock,
			&article.Created,
			&article.Updated,
			&article.Enabled,
		)
		if err != nil {
			log.Get(deps...).Error(err)
			return nil, err
		}
		articles = append(articles, &article)
	}

	if rows.Err() != nil {
		log.Get(deps...).Error(rows.Err())
		return nil, rows.Err()
	}

	return articles, nil
}

func findById(articleId string, deps ...interface{}) (result *Article, err error) {
	conn, err := db.GetPostgresClient(deps...)
	if err != nil {
		log.Get(deps...).Error(err)

		return nil, err
	}

	var article Article
	err = conn.QueryRow(context.Background(), "SELECT id, name, description, image, price, stock, created, updated, enabled FROM Articles WHERE id=$1", articleId).Scan(
		&article.ID,
		&article.Name,
		&article.Description,
		&article.Image,
		&article.Price,
		&article.Stock,
		&article.Created,
		&article.Updated,
		&article.Enabled,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errs.NotFound
		}
		return nil, err
	}

	return &article, nil
}

func insert(article *Article, deps ...interface{}) (err error) {
	if err := article.validateSchema(); err != nil {
		log.Get(deps...).Error(err)
		return err
	}

	conn, err := db.GetPostgresClient(deps...)
	if err != nil {
		log.Get(deps...).Error(err)
		return err
	}

	query := `
        INSERT INTO Articles (id, name, description, image, price, stock, created, updated, enabled)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
    `

	_, err = conn.Exec(context.Background(), query,
		article.ID,
		article.Name,
		article.Description,
		article.Image,
		article.Price,
		article.Stock,
		article.Created,
		article.Updated,
		article.Enabled,
	)
	if err != nil {
		log.Get(deps...).Error(err)
		return err
	}

	return nil
}

// disable Deshabilita el articulo para que no se pueda usar mas
func Disable(articleId string, deps ...interface{}) (err error) {
	conn, err := db.GetPostgresClient(deps...)
	if err != nil {
		log.Get(deps...).Error(err)
		return err
	}

	query := `
        UPDATE Articles
        SET enabled = $1
        WHERE id = $2
    `

	_, err = conn.Exec(context.Background(), query, false, articleId)
	if err != nil {
		log.Get(deps...).Error(err)
		return err
	}

	return nil
}

// Actualiza la descripción de un articulo.
func update(articleId string, article UpdateArticleData, deps ...interface{}) (err error) {
	conn, err := db.GetPostgresClient(deps...)
	if err != nil {
		log.Get(deps...).Error(err)
		return err
	}

	query := `
        UPDATE Articles
        SET name = $1, description = $2, image = $3, price = $4, stock = $5, updated = NOW()
        WHERE id = $7
    `

	_, err = conn.Exec(context.Background(), query,
		article.Name,
		article.Description,
		article.Image,
		article.Price,
		article.Stock,
		articleId,
	)
	if err != nil {
		log.Get(deps...).Error(err)
		return err
	}

	return nil
}

func DecrementStock(articleId string, amount int, deps ...interface{}) (err error) {
	conn, err := db.GetPostgresClient(deps...)
	if err != nil {
		log.Get(deps...).Error(err)
		return err
	}

	query := `
        UPDATE Articles
        SET stock = stock - $1
        WHERE id = $2
    `

	_, err = conn.Exec(context.Background(), query, amount, articleId)
	if err != nil {
		log.Get(deps...).Error(err)
		return err
	}

	return nil
}

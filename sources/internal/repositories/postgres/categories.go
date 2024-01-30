package postgresrepository

import (
	"context"
	"database/sql"
	"fmt"
	postgresclient "shtem-web/sources/internal/clients/postgres"
	"shtem-web/sources/internal/core/domain"
)

var categoriesTableName = "categories"

var categoriesTableComponents = struct {
	c_id        string
	name        string
	description string
}{
	c_id:        categoriesTableName + ".c_id",
	name:        categoriesTableName + ".name",
	description: categoriesTableName + ".description",
}

type categoriesDB struct {
	ctx context.Context
	db  *postgresclient.PostgresDB
}

func (q *categoriesDB) GetCategories() ([]*domain.Category, domain.Error) {
	var categories []*domain.Category

	// FIND DISTINCT SHTEMARAN NAMES
	query := fmt.Sprintf(`
		SELECT %s, %s, %s
		FROM %s`,
		categoriesTableComponents.c_id,
		categoriesTableComponents.name,
		categoriesTableComponents.description,
		categoriesTableName, // TABLE NAME
	)

	rows, err := q.db.Query(q.ctx, query)
	if err != nil {
		return nil, domain.NewError().SetError(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int64
		var name, description sql.NullString

		if err := rows.Scan(
			&id,
			&name,
			&description,
		); err != nil {
			return nil, domain.NewError().SetError(err)
		}

		categories = append(categories, &domain.Category{
			C_id:        id,
			Name:        name.String,
			Description: description.String,
		})
	}

	if err := rows.Err(); err != nil {
		return nil, domain.NewError().SetError(err)
	}

	return categories, nil
}

func (q *categoriesDB) GetCategoriesWithShtems() (domain.Categories, domain.Error) {
	categories := make(domain.Categories)

	// FIND DISTINCT SHTEMARAN NAMES
	query := fmt.Sprintf(`
		SELECT
			COUNT(%s) AS arraysCount,
			%s AS category,
			%s AS c_description,
			ARRAY_AGG(%s) AS names,
			ARRAY_AGG(%s) AS descriptions,
			ARRAY_AGG(%s) AS link_names,
			ARRAY_AGG(%s) AS images,
			ARRAY_AGG(%s) AS authors
		FROM %s
		LEFT JOIN %s
		ON %s = %s
		GROUP BY %s;`,
		// ARRAY_ARG
		categoriesTableComponents.name,
		categoriesTableComponents.name,
		categoriesTableComponents.description,
		shtemsTableComponents.name,
		shtemsTableComponents.description,
		shtemsTableComponents.link_name,
		shtemsTableComponents.image,
		shtemsTableComponents.author,
		// FROM TABLE NAME
		categoriesTableName,
		// LEFT JOIN TABLE NAME
		shtemsTableName,
		// ON
		categoriesTableComponents.c_id,
		shtemsTableComponents.category,
		// GROUP BY
		categoriesTableComponents.c_id,
	)

	rows, err := q.db.Query(q.ctx, query)
	if err != nil {
		return nil, domain.NewError().SetError(err)
	}
	defer rows.Close()

	for rows.Next() {
		var arraysCount int
		var category string
		var c_description sql.NullString
		var names, descriptions, link_names, images, authors []sql.NullString

		if err := rows.Scan(
			&arraysCount,
			&category,
			&c_description,
			&names,
			&descriptions,
			&link_names,
			&images,
			&authors,
		); err != nil {
			return nil, domain.NewError().SetError(err)
		}

		for i := 0; i < arraysCount; i++ {
			c := domain.Category{
				Name:        category,
				Description: c_description.String,
			}
			s := &domain.Shtemaran{
				Name:        names[i].String,
				Description: descriptions[i].String,
				LinkName:    link_names[i].String,
				Image:       images[i].String,
				Author:      authors[i].String,
			}

			categories[c] = append(categories[c], s)
		}
	}

	if err := rows.Err(); err != nil {
		return nil, domain.NewError().SetError(err)
	}

	return categories, nil
}

func NewCategoriesDB(ctx context.Context, db *postgresclient.PostgresDB) *categoriesDB {
	return &categoriesDB{ctx, db}
}

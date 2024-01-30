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

func NewCategoriesDB(ctx context.Context, db *postgresclient.PostgresDB) *categoriesDB {
	return &categoriesDB{ctx, db}
}

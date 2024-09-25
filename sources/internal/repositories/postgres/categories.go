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
	link_name   string
	score       string
}{
	c_id:        categoriesTableName + ".c_id",
	name:        categoriesTableName + ".name",
	description: categoriesTableName + ".description",
	link_name:   categoriesTableName + ".link_name",
	score:       categoriesTableName + ".score",
}

type categoriesDB struct {
	ctx context.Context
	db  *postgresclient.PostgresDB
}

// Helper function to execute category select queries
func (q *categoriesDB) SelectCategoryFields(query string, args ...interface{}) ([]*domain.Category, domain.Error) {
	var categories []*domain.Category
	rows, err := q.db.Query(q.ctx, query, args...)
	if err != nil {
		return nil, domain.NewError().SetError(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int64
		var name, description, link_name sql.NullString
		if err := rows.Scan(&id, &name, &description, &link_name); err != nil {
			return nil, domain.NewError().SetError(err)
		}
		categories = append(categories, &domain.Category{
			C_id:        id,
			Name:        name.String,
			Description: description.String,
			LinkName:    link_name.String,
		})
	}

	if err := rows.Err(); err != nil {
		return nil, domain.NewError().SetError(err)
	}

	if len(categories) == 0 {
		return nil, domain.ErrNoRows
	}

	return categories, nil
}

// Function to get all categories
func (q *categoriesDB) GetCategories() ([]*domain.Category, domain.Error) {
	query := fmt.Sprintf(`
		SELECT %s, %s, %s, %s
		FROM %s
		ORDER BY %s DESC`,
		categoriesTableComponents.c_id,
		categoriesTableComponents.name,
		categoriesTableComponents.description,
		categoriesTableComponents.link_name,
		categoriesTableName,
		categoriesTableComponents.score,
	)

	return q.SelectCategoryFields(query)
}

// Function to get a category by its link name
func (q *categoriesDB) GetCategoryByLinkName(linkName string) (*domain.Category, domain.Error) {
	query := fmt.Sprintf(`
		SELECT %s, %s, %s, %s
		FROM %s
		WHERE %s = $1
		LIMIT 1`,
		categoriesTableComponents.c_id,
		categoriesTableComponents.name,
		categoriesTableComponents.description,
		categoriesTableComponents.link_name,
		categoriesTableName,
		categoriesTableComponents.link_name,
	)

	categories, err := q.SelectCategoryFields(query, linkName)
	if err != nil || len(categories) == 0 {
		return nil, err
	}

	return categories[0], nil
}

// Function to get a category by its ID
func (q *categoriesDB) GetCategoryByID(id int64) (*domain.Category, domain.Error) {
	query := fmt.Sprintf(`
		SELECT %s, %s, %s, %s
		FROM %s
		WHERE %s = $1
		LIMIT 1`,
		categoriesTableComponents.c_id,
		categoriesTableComponents.name,
		categoriesTableComponents.description,
		categoriesTableComponents.link_name,
		categoriesTableName,
		categoriesTableComponents.c_id,
	)

	categories, err := q.SelectCategoryFields(query, id)
	if err != nil || len(categories) == 0 {
		return nil, err
	}

	return categories[0], nil
}

// Constructor function for creating a new categoriesDB instance
func NewCategoriesDB(ctx context.Context, db *postgresclient.PostgresDB) *categoriesDB {
	return &categoriesDB{ctx, db}
}

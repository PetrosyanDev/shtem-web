// Erik Petrosyan ©
package postgresrepository

import (
	"context"
	"database/sql"
	"fmt"
	postgresclient "shtem-web/sources/internal/clients/postgres"
	"shtem-web/sources/internal/core/domain"
)

var shtemsTableComponents = struct {
	id          string
	name        string
	description string
	author      string
	link_name   string
	image       string
	pdf         string
}{
	id:          "id",
	name:        "name",
	description: "description",
	author:      "author",
	link_name:   "link_name",
	image:       "image",
	pdf:         "pdf",
}

var shtemsTableName = "shtems"

type shtemsDB struct {
	ctx context.Context
	db  *postgresclient.PostgresDB
}

func (q *shtemsDB) GetShtemByLinkName(name string) (*domain.Shtemaran, domain.Error) {

	var result *domain.Shtemaran

	query := fmt.Sprintf("SELECT %s, %s, %s, %s, %s, %s FROM %s WHERE %s=$1 LIMIT 1",
		shtemsTableComponents.name,
		shtemsTableComponents.description,
		shtemsTableComponents.author,
		shtemsTableComponents.link_name,
		shtemsTableComponents.image,
		shtemsTableComponents.pdf,
		shtemsTableName,                 // TABLE NAME
		shtemsTableComponents.link_name, // LINK NAME
	)

	rows, err := q.db.Query(q.ctx, query, name)
	if err != nil {
		return nil, domain.NewError().SetError(err)
	}
	defer rows.Close()

	if rows.Next() {
		var name, description, author, linkName, image, pdf sql.NullString

		if err := rows.Scan(
			&name,
			&description,
			&author,
			&linkName,
			&image,
			&pdf,
		); err != nil {
			return nil, domain.NewError().SetError(err)
		}

		result = &domain.Shtemaran{
			Name:        name.String,
			Description: description.String,
			Author:      author.String,
			LinkName:    linkName.String,
			Image:       image.String,
			PDF:         pdf.String,
		}
	}

	return result, nil
}

func (q *shtemsDB) GetShtemNames() ([]*domain.Shtemaran, domain.Error) {
	var shtemarans []*domain.Shtemaran

	// FIND DISTINCT SHTEMARAN NAMES
	query := fmt.Sprintf("SELECT %s, %s, %s, %s, %s, %s FROM %s",
		shtemsTableComponents.name,
		shtemsTableComponents.description,
		shtemsTableComponents.author,
		shtemsTableComponents.link_name,
		shtemsTableComponents.image,
		shtemsTableComponents.pdf,
		shtemsTableName, // TABLE NAME
	)

	rows, err := q.db.Query(q.ctx, query)
	if err != nil {
		return nil, domain.NewError().SetError(err)
	}
	defer rows.Close()

	for rows.Next() {
		var name, description, author, linkName, image, pdf sql.NullString

		if err := rows.Scan(
			&name,
			&author,
			&description,
			&linkName,
			&image,
			&pdf,
		); err != nil {
			return nil, domain.NewError().SetError(err)
		}

		shtemarans = append(shtemarans, &domain.Shtemaran{
			Name:        name.String,
			Description: description.String,
			Author:      author.String,
			LinkName:    linkName.String,
			Image:       image.String,
			PDF:         pdf.String,
		})
	}

	if err := rows.Err(); err != nil {
		return nil, domain.NewError().SetError(err)
	}

	return shtemarans, nil
}

func NewShtemsDB(ctx context.Context, db *postgresclient.PostgresDB) *shtemsDB {
	return &shtemsDB{ctx, db}
}

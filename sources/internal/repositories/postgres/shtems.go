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
	link_name   string
	image       string
}{
	id:          "id",
	name:        "name",
	description: "description",
	link_name:   "link_name",
	image:       "image",
}

var shtemsTableName = "shtems"

type shtemsDB struct {
	ctx context.Context
	db  *postgresclient.PostgresDB
}

func (q *shtemsDB) GetShtemByLinkName(name string) (*domain.Shtemaran, domain.Error) {

	var result *domain.Shtemaran

	query := fmt.Sprintf("SELECT %s, %s, %s, %s FROM %s WHERE %s=$1 LIMIT 1",
		shtemsTableComponents.name,
		shtemsTableComponents.description,
		shtemsTableComponents.link_name,
		shtemsTableComponents.image,
		shtemsTableName,                 // TABLE NAME
		shtemsTableComponents.link_name, // LINK NAME
	)

	rows, err := q.db.Query(q.ctx, query, name)
	if err != nil {
		return nil, domain.NewError().SetError(err)
	}
	defer rows.Close()

	if rows.Next() {
		var (
			name        string
			description string
			linkName    string
			image       string
		)
		if err := rows.Scan(&name, &description, &linkName, &image); err != nil {
			return nil, domain.NewError().SetError(err)
		}

		result = &domain.Shtemaran{
			Name:        name,
			Description: description,
			LinkName:    linkName,
			Image:       image,
		}
	}

	return result, nil
}

func (q *shtemsDB) GetShtemNames() ([]*domain.Shtemaran, domain.Error) {
	var shtemarans []*domain.Shtemaran

	// FIND DISTINCT SHTEMARAN NAMES
	query := fmt.Sprintf("SELECT %s, %s, %s, %s FROM %s",
		shtemsTableComponents.name,
		shtemsTableComponents.description,
		shtemsTableComponents.link_name,
		shtemsTableComponents.image,
		shtemsTableName, // TABLE NAME
	)

	rows, err := q.db.Query(q.ctx, query)
	if err != nil {
		return nil, domain.NewError().SetError(err)
	}
	defer rows.Close()

	for rows.Next() {
		var shtem domain.Shtemaran
		var name, description, linkName, image sql.NullString

		if err := rows.Scan(
			&name,
			&description,
			&linkName,
			&image,
		); err != nil {
			return nil, domain.NewError().SetError(err)
		}

		shtem.Name = name.String
		shtem.Description = description.String
		shtem.LinkName = linkName.String
		shtem.Image = image.String

		shtemarans = append(shtemarans, &shtem)
	}

	if err := rows.Err(); err != nil {
		return nil, domain.NewError().SetError(err)
	}

	return shtemarans, nil
}

func NewShtemsDB(ctx context.Context, db *postgresclient.PostgresDB) *shtemsDB {
	return &shtemsDB{ctx, db}
}

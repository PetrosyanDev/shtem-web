// Erik Petrosyan ©
package postgresrepository

import (
	"context"
	"fmt"
	postgresclient "shtem-web/sources/internal/clients/postgres"
	"shtem-web/sources/internal/core/domain"
)

var shtemsTableComponents = struct {
	id          string
	name        string
	description string
	link_name   string
}{
	id:          "id",
	name:        "name",
	description: "description",
	link_name:   "link_name",
}

var shtemsTableName = "shtems"

type shtemsDB struct {
	ctx context.Context
	db  *postgresclient.PostgresDB
}

func (q *shtemsDB) GetShtemNames() ([]*domain.Shtemaran, domain.Error) {
	var shtemarans []*domain.Shtemaran

	// FIND DISTINCT SHTEMARAN NAMES
	query := fmt.Sprintf("SELECT %s, %s, %s FROM %s",
		shtemsTableComponents.name,
		shtemsTableComponents.description,
		shtemsTableComponents.link_name,
		shtemsTableName, // TABLE NAME
	)

	rows, err := q.db.Query(q.ctx, query)
	if err != nil {
		return nil, domain.NewError().SetError(err)
	}
	defer rows.Close()

	for rows.Next() {
		var shtem domain.Shtemaran
		if err := rows.Scan(
			&shtem.Name,
			&shtem.Description,
			&shtem.LinkName,
		); err != nil {
			return nil, domain.NewError().SetError(err)
		}
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

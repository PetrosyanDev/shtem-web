// Erik Petrosyan ©
package postgresrepository

import (
	"context"
	"fmt"
	postgresclient "shtem-web/sources/internal/clients/postgres"
	"shtem-web/sources/internal/core/domain"
)

var questionsTableComponents = struct {
	q_id      string
	shtemaran string
	bajin     string
	mas       string
	q_number  string
	text      string
	options   string
	answers   string
}{
	q_id:      "q_id",
	shtemaran: "shtemaran",
	bajin:     "bajin",
	mas:       "mas",
	q_number:  "q_number",
	text:      "text",
	options:   "options",
	answers:   "answers",
}

var questionsTableName = "questions"

type questionsDB struct {
	ctx context.Context
	db  *postgresclient.PostgresDB
}

func (q *questionsDB) GetShtemNames() ([]string, domain.Error) {
	var shtemaranNames []string

	// FIND DISTINCT SHTEMARAN NAMES
	query := fmt.Sprintf("SELECT DISTINCT %s FROM %s",
		questionsTableComponents.shtemaran,
		questionsTableName, // TABLE NAME
	)

	rows, err := q.db.Query(q.ctx, query)
	if err != nil {
		return nil, domain.NewError().SetError(err)
	}
	defer rows.Close()

	for rows.Next() {
		var shtemaranName string
		if err := rows.Scan(&shtemaranName); err != nil {
			return nil, domain.NewError().SetError(err)
		}
		shtemaranNames = append(shtemaranNames, shtemaranName)
	}

	if err := rows.Err(); err != nil {
		return nil, domain.NewError().SetError(err)
	}

	return shtemaranNames, nil
}

func NewQuestionsDB(ctx context.Context, db *postgresclient.PostgresDB) *questionsDB {
	return &questionsDB{ctx, db}
}

// Erik Petrosyan ©
package postgresrepository

import (
	"context"
	"database/sql"
	"fmt"
	postgresclient "shtem-web/sources/internal/clients/postgres"
	"shtem-web/sources/internal/core/domain"
	"strings"
)

var emailsTableName = "emails"

type emailsTable struct {
	id        string
	email     string
	createdAt string
}

var emailsTableComponents = emailsTable{
	id:        emailsTableName + ".id",
	email:     emailsTableName + ".email",
	createdAt: emailsTableName + ".created_at",
}

type emailsDB struct {
	ctx context.Context
	db  *postgresclient.PostgresDB
}

// InsertEmail inserts a new email record into the emails table
func (q *emailsDB) InsertEmail(email string) domain.Error {
	query := fmt.Sprintf(`
		INSERT INTO %s (email) 
		VALUES ($1)`,
		emailsTableName,
	)

	_, err := q.db.Exec(q.ctx, query, email)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return domain.NewError().SetMessage("Email already exists")
		}
		return domain.NewError().SetError(err)
	}

	return nil
}

// GetAllEmails retrieves all email records from the emails table
func (q *emailsDB) GetAllEmails() ([]*domain.Email, domain.Error) {
	var emails []*domain.Email

	query := fmt.Sprintf(`
		SELECT %s, %s, %s 
		FROM %s
		ORDER BY %s DESC`,
		emailsTableComponents.id,
		emailsTableComponents.email,
		emailsTableComponents.createdAt,
		emailsTableName,
		emailsTableComponents.createdAt,
	)

	rows, err := q.db.Query(q.ctx, query)
	if err != nil {
		return nil, domain.NewError().SetError(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int64
		var email string
		var createdAt sql.NullTime

		if err := rows.Scan(
			&id,
			&email,
			&createdAt,
		); err != nil {
			return nil, domain.NewError().SetError(err)
		}

		emails = append(emails, &domain.Email{
			Id:        id,
			Email:     email,
			CreatedAt: createdAt.Time,
		})
	}

	if err := rows.Err(); err != nil {
		return nil, domain.NewError().SetError(err)
	}

	return emails, nil
}

// NewEmailsDB creates a new instance of the emailsDB struct
func NewEmailsDB(ctx context.Context, db *postgresclient.PostgresDB) *emailsDB {
	return &emailsDB{ctx, db}
}

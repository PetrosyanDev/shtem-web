// Erik Petrosyan ©
package postgresrepository

import (
	"context"
	"fmt"
	postgreclient "shtem-web/sources/internal/clients/postgres"
	"shtem-web/sources/internal/core/domain"

	"github.com/jackc/pgx/v5"
)

var questionsTableName = "questions"

var questionsTableComponents = struct {
	q_id     string
	bajin    string
	mas      string
	q_number string
	text     string
	options  string
	answers  string
	shtem_id string
}{
	q_id:     questionsTableName + "q_id",
	bajin:    questionsTableName + "bajin",
	mas:      questionsTableName + "mas",
	q_number: questionsTableName + "q_number",
	text:     questionsTableName + "text",
	options:  questionsTableName + "options",
	answers:  questionsTableName + "answers",
	shtem_id: questionsTableName + "shtem_id",
}

type questionsDB struct {
	ctx context.Context
	db  *postgreclient.PostgresDB
}

// CREATE!
// CREATE!
// CREATE!
func (q *questionsDB) Create(question *domain.Question) domain.Error {

	query := fmt.Sprintf(`
		INSERT INTO %s (%s,%s,%s,%s,%s,%s,%s) 
		VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		questionsTableName, // TABLE NAME
		questionsTableComponents.bajin,
		questionsTableComponents.mas,
		questionsTableComponents.q_number,
		questionsTableComponents.text,
		questionsTableComponents.options,
		questionsTableComponents.answers,
		questionsTableComponents.shtem_id,
	)
	_, err := q.db.Exec(q.ctx, query,
		question.Bajin,
		question.Mas,
		question.Q_number,
		question.Text,
		question.Options,
		question.Answers,
		question.ShtemId,
	)
	if err != nil {
		return domain.NewError().SetError(err)
	}

	return nil
}

// UPDATE!
// UPDATE!
// UPDATE!
func (q *questionsDB) Update(question *domain.Question) domain.Error {

	query := fmt.Sprintf(`
		UPDATE %s 
		SET %s=$1, %s=$2, %s=$3, %s=$4, %s=$5, %s=$6, %s=$7 
		WHERE %s=$8`,
		questionsTableName, // TABLE NAME
		questionsTableComponents.bajin,
		questionsTableComponents.mas,
		questionsTableComponents.q_number,
		questionsTableComponents.text,
		questionsTableComponents.options,
		questionsTableComponents.answers,
		questionsTableComponents.shtem_id,
		questionsTableComponents.q_id, // for identifying the question to update
	)
	_, err := q.db.Exec(q.ctx, query,
		question.Bajin,
		question.Mas,
		question.Q_number,
		question.Text,
		question.Options,
		question.Answers,
		question.ShtemId,
		question.Q_id, // for identifying the question to update
	)
	if err != nil {
		return domain.NewError().SetError(err)
	}
	return nil
}

// DELETE!
// DELETE!
// DELETE!
func (q *questionsDB) Delete(id int64) domain.Error {
	// DELETE!
	query := fmt.Sprintf(`
		DELETE FROM %s 
		WHERE %s=$1`,
		questionsTableName,
		questionsTableComponents.q_id,
	)
	_, err := q.db.Exec(q.ctx, query,
		id,
	)
	if err != nil {
		return domain.NewError().SetError(err)
	}

	return nil
}

// FINDQUESTION
// FINDQUESTION
// FINDQUESTION
func (q *questionsDB) FindQuestion(question *domain.Question) (*domain.Question, domain.Error) {

	var result domain.Question

	// FIND!
	query := fmt.Sprintf(`
		SELECT %s, %s, %s, %s, %s, %s, %s 
		FROM %s 
		WHERE %s=$1 AND %s=$2 AND %s=$3 AND %s=$4`,
		// SELECT
		questionsTableComponents.bajin,
		questionsTableComponents.mas,
		questionsTableComponents.q_number,
		questionsTableComponents.text,
		questionsTableComponents.options,
		questionsTableComponents.answers,
		questionsTableComponents.shtem_id,
		// FROM
		questionsTableName,
		// WHERE
		questionsTableComponents.shtem_id, // shtems
		questionsTableComponents.bajin,    // bajin
		questionsTableComponents.mas,      // mas
		questionsTableComponents.q_number, // q_number
	)
	err := q.db.QueryRow(q.ctx, query,
		// WHERE
		question.ShtemId,
		question.Bajin,
		question.Mas,
		question.Q_number).
		Scan(
			// SELECTED
			&result.Bajin,
			&result.Mas,
			&result.Q_number,
			&result.Text,
			&result.Options,
			&result.Answers,
			&result.ShtemId,
		)
	if err == pgx.ErrNoRows {
		return nil, domain.ErrNoRows
	} else if err != nil {
		return nil, domain.NewError().SetError(err)
	}

	return &result, nil
}

// FINDBAJIN
// FINDBAJIN
// FINDBAJIN
func (q *questionsDB) FindBajin(question *domain.Question) ([]*domain.Question, domain.Error) {

	// FIND!
	query := fmt.Sprintf(`
		SELECT %s, %s, %s, %s, %s, %s, %s 
		FROM %s 
		JOIN %s
		ON %s=%s
		WHERE %s=$2 AND %s=$3
		ORDER BY %s, %s`,
		// SELECT
		questionsTableComponents.bajin,
		questionsTableComponents.mas,
		questionsTableComponents.q_number,
		questionsTableComponents.text,
		questionsTableComponents.options,
		questionsTableComponents.answers,
		questionsTableComponents.shtem_id,
		// FROM
		questionsTableName,
		// JOIN
		shtemsTableName,
		// ON
		shtemsTableComponents.id,
		questionsTableComponents.shtem_id,
		// WHERE
		questionsTableComponents.shtem_id,
		questionsTableComponents.bajin,
		// ORDER BY
		questionsTableComponents.mas,
		questionsTableComponents.q_number,
	)
	rows, err := q.db.Query(q.ctx, query,
		question.ShtemId,
		question.Bajin,
	)
	if err == pgx.ErrNoRows {
		return nil, domain.ErrNoRows
	} else if err != nil {
		return nil, domain.NewError().SetError(err)
	}
	defer rows.Close()

	var outputQuestions []*domain.Question

	for rows.Next() {
		var result domain.Question
		// Scan the row data into the result struct
		if err := rows.Scan(
			&result.Bajin,
			&result.Mas,
			&result.Q_number,
			&result.Text,
			&result.Options,
			&result.Answers,
			&result.ShtemId,
		); err != nil {
			return nil, domain.ErrBadRequest
		}

		outputQuestions = append(outputQuestions, &result)
	}

	return outputQuestions, nil
}

// FINDBYID!
// FINDBYID!
// FINDBYID!
func (q *questionsDB) FindByID(id int) (*domain.Question, domain.Error) {

	var result domain.Question

	query := fmt.Sprintf(`
		SELECT %s, %s, %s, %s, %s, %s, %s, %s 
		FROM %s 
		WHERE %s=$1`,
		// SELECT
		questionsTableComponents.q_id,
		questionsTableComponents.bajin,
		questionsTableComponents.mas,
		questionsTableComponents.q_number,
		questionsTableComponents.text,
		questionsTableComponents.options,
		questionsTableComponents.answers,
		questionsTableComponents.shtem_id,
		// FROM
		questionsTableName,
		// WHERE
		questionsTableComponents.q_id,
	)
	err := q.db.QueryRow(q.ctx, query, id).
		Scan(
			&result.Q_id,
			&result.Bajin,
			&result.Mas,
			&result.Q_number,
			&result.Text,
			&result.Options,
			&result.Answers,
			&result.ShtemId,
		)
	if err == pgx.ErrNoRows {
		return nil, domain.ErrNoRows
	} else if err != nil {
		return nil, domain.NewError().SetError(err)
	}

	return &result, nil
}

func NewQuestionsDB(ctx context.Context, db *postgreclient.PostgresDB) *questionsDB {
	return &questionsDB{ctx, db}
}

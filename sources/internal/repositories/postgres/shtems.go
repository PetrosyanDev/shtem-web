// Erik Petrosyan ©
package postgresrepository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	postgresclient "shtem-web/sources/internal/clients/postgres"
	"shtem-web/sources/internal/core/domain"
)

var shtemsTableName = "shtems"

type shtemsTable struct {
	id          string
	name        string
	description string
	author      string
	link_name   string
	image       string
	pdf         string
	category    string
	keywords    string
	has_quiz    string
	has_pdf     string
}

var shtemsTableComponents = shtemsTable{
	id:          shtemsTableName + ".id",
	name:        shtemsTableName + ".name",
	description: shtemsTableName + ".description",
	author:      shtemsTableName + ".author",
	link_name:   shtemsTableName + ".link_name",
	image:       shtemsTableName + ".image",
	pdf:         shtemsTableName + ".pdf",
	keywords:    shtemsTableName + ".keywords",
	category:    shtemsTableName + ".category",
	has_quiz:    shtemsTableName + ".has_quiz",
	has_pdf:     shtemsTableName + ".has_pdf",
}

// var shtemsTableComponentsNon = shtemsTable{
// 	id:          "id",
// 	name:        "name",
// 	description: "description",
// 	author:      "author",
// 	link_name:   "link_name",
// 	image:       "image",
// 	pdf:         "pdf",
// 	keywords:    "keywords",
// 	category:    "category",
// 	has_quiz:    "has_quiz",
// 	has_pdf:     "has_pdf",
// }

type shtemsDB struct {
	ctx context.Context
	db  *postgresclient.PostgresDB
}

func (q *shtemsDB) GetShtemByLinkName(name string) (*domain.Shtemaran, domain.Error) {

	var result *domain.Shtemaran

	query := fmt.Sprintf(`
		SELECT %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s
		FROM %s 
		WHERE %s=$1
		LIMIT 1`,
		shtemsTableComponents.id,
		shtemsTableComponents.name,
		shtemsTableComponents.description,
		shtemsTableComponents.author,
		shtemsTableComponents.link_name,
		shtemsTableComponents.image,
		shtemsTableComponents.pdf,
		shtemsTableComponents.keywords,
		shtemsTableComponents.category,
		shtemsTableComponents.has_quiz,
		shtemsTableComponents.has_pdf,
		shtemsTableName,                 // TABLE NAME
		shtemsTableComponents.link_name, // LINK NAME
	)

	rows, err := q.db.Query(q.ctx, query, name)
	if err != nil {
		return nil, domain.NewError().SetError(err)
	}
	defer rows.Close()

	if rows.Next() {
		var id, category int64
		var name, description, author, linkName, image, pdf sql.NullString
		var keywords []string
		var has_quiz, has_pdf bool

		if err := rows.Scan(
			&id,
			&name,
			&description,
			&author,
			&linkName,
			&image,
			&pdf,
			&keywords,
			&category,
			&has_quiz,
			&has_pdf,
		); err != nil {
			return nil, domain.NewError().SetError(err)
		}

		result = &domain.Shtemaran{
			Id:          id,
			Name:        name.String,
			Description: description.String,
			Author:      author.String,
			LinkName:    linkName.String,
			Image:       image.String,
			PDF:         pdf.String,
			Keywords:    keywords,
			Category:    category,
			HasQuiz:     has_quiz,
			HasPDF:      has_pdf,
		}
	}

	return result, nil
}

func (q *shtemsDB) GetShtems() ([]*domain.Shtemaran, domain.Error) {
	var shtemarans []*domain.Shtemaran

	// FIND DISTINCT SHTEMARAN NAMES
	query := fmt.Sprintf(`
		SELECT %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s 
		FROM %s
		ORDER BY %s`,
		shtemsTableComponents.id,
		shtemsTableComponents.name,
		shtemsTableComponents.description,
		shtemsTableComponents.author,
		shtemsTableComponents.link_name,
		shtemsTableComponents.image,
		shtemsTableComponents.pdf,
		shtemsTableComponents.keywords,
		shtemsTableComponents.category,
		shtemsTableComponents.has_quiz,
		shtemsTableComponents.has_pdf,
		shtemsTableName, // TABLE NAME
		shtemsTableComponents.link_name,
	)

	rows, err := q.db.Query(q.ctx, query)
	if err != nil {
		return nil, domain.NewError().SetError(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, category int64
		var name, description, author, linkName, image, pdf sql.NullString
		var keywords []string
		var has_quiz, has_pdf bool

		if err := rows.Scan(
			&id,
			&name,
			&description,
			&author,
			&linkName,
			&image,
			&pdf,
			&keywords,
			&category,
			&has_quiz,
			&has_pdf,
		); err != nil {
			return nil, domain.NewError().SetError(err)
		}

		shtemarans = append(shtemarans, &domain.Shtemaran{
			Id:          id,
			Name:        name.String,
			Description: description.String,
			Author:      author.String,
			LinkName:    linkName.String,
			Image:       image.String,
			PDF:         pdf.String,
			Keywords:    keywords,
			Category:    category,
			HasQuiz:     has_quiz,
			HasPDF:      has_pdf,
		})
	}

	if err := rows.Err(); err != nil {
		return nil, domain.NewError().SetError(err)
	}

	return shtemarans, nil
}

func (q *shtemsDB) GetShtemLinkNames() ([]string, domain.Error) {
	var linkNames []string

	// FIND DISTINCT SHTEMARAN NAMES
	query := fmt.Sprintf(`
		SELECT DISTINCT %s
		FROM %s
		ORDER BY %s`,
		shtemsTableComponents.link_name,
		shtemsTableName, // TABLE NAME
		shtemsTableComponents.link_name,
	)

	rows, err := q.db.Query(q.ctx, query)
	if err != nil {
		return nil, domain.NewError().SetError(err)
	}
	defer rows.Close()

	for rows.Next() {
		var linkName sql.NullString

		if err := rows.Scan(
			&linkName,
		); err != nil {
			return nil, domain.NewError().SetError(err)
		}

		linkNames = append(linkNames, linkName.String)
	}

	if err := rows.Err(); err != nil {
		return nil, domain.NewError().SetError(err)
	}

	return linkNames, nil
}

func (q *shtemsDB) GetShtemsByCategoryId(c_id int64) ([]*domain.Shtemaran, domain.Error) {
	var shtemarans []*domain.Shtemaran

	// FIND DISTINCT SHTEMARAN NAMES
	query := fmt.Sprintf(`
		SELECT %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s
		FROM %s
		WHERE %s = $1
		ORDER BY %s DESC`,
		shtemsTableComponents.id,
		shtemsTableComponents.name,
		shtemsTableComponents.description,
		shtemsTableComponents.author,
		shtemsTableComponents.link_name,
		shtemsTableComponents.image,
		shtemsTableComponents.pdf,
		shtemsTableComponents.keywords,
		shtemsTableComponents.category,
		shtemsTableComponents.has_quiz,
		shtemsTableComponents.has_pdf,
		shtemsTableName,                // TABLE NAME
		shtemsTableComponents.category, // MATCH
		shtemsTableComponents.name,     // ORDER BY
	)

	rows, err := q.db.Query(q.ctx, query, c_id)
	if err != nil {
		return nil, domain.NewError().SetError(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, category int64
		var name, description, author, linkName, image, pdf sql.NullString
		var keywords []string
		var has_quiz, has_pdf bool

		if err := rows.Scan(
			&id,
			&name,
			&description,
			&author,
			&linkName,
			&image,
			&pdf,
			&keywords,
			&category,
			&has_quiz,
			&has_pdf,
		); err != nil {
			return nil, domain.NewError().SetError(err)
		}

		log.Println(linkName)

		shtemarans = append(shtemarans, &domain.Shtemaran{
			Id:          id,
			Name:        name.String,
			Description: description.String,
			Author:      author.String,
			LinkName:    linkName.String,
			Image:       image.String,
			PDF:         pdf.String,
			Keywords:    keywords,
			Category:    category,
			HasQuiz:     has_quiz,
			HasPDF:      has_pdf,
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

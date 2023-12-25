package repositories

import (
	"shtem-web/sources/internal/core/domain"
	postgresrepository "shtem-web/sources/internal/repositories/postgres"
)

type questionsRepository struct {
	db postgresrepository.QuestionsDB
}

func (p *questionsRepository) GetShtemNames() ([]string, domain.Error) {
	return p.db.GetShtemNames()
}

func NewQuestionsRepository(db postgresrepository.QuestionsDB) *questionsRepository {
	return &questionsRepository{db}
}

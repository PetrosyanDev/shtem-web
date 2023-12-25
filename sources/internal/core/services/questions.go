package services

import (
	"shtem-web/sources/internal/core/domain"
	"shtem-web/sources/internal/repositories"
)

type questionsService struct {
	questionsRepository repositories.QuestionsRepository
}

func (q *questionsService) GetShtemNames() ([]string, domain.Error) {
	return q.questionsRepository.GetShtemNames()
}

func NewQuestionsService(questionsRepository repositories.QuestionsRepository) *questionsService {
	return &questionsService{questionsRepository}
}

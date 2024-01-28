package services

import (
	"shtem-web/sources/internal/core/domain"
	"shtem-web/sources/internal/repositories"
)

type shtemsService struct {
	shtemsRepository repositories.ShtemsRepository
}

func (q *shtemsService) GetShtemNames() ([]*domain.Shtemaran, domain.Error) {
	return q.shtemsRepository.GetShtemNames()
}
func (q *shtemsService) GetShtemByLinkName(name string) (*domain.Shtemaran, domain.Error) {
	return q.shtemsRepository.GetShtemByLinkName(name)
}

func NewShtemsService(shtemsRepository repositories.ShtemsRepository) *shtemsService {
	return &shtemsService{shtemsRepository}
}

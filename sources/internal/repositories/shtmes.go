package repositories

import (
	"shtem-web/sources/internal/core/domain"
	postgresrepository "shtem-web/sources/internal/repositories/postgres"
)

type shtemsRepository struct {
	db postgresrepository.ShtemsDB
}

func (p *shtemsRepository) GetShtemNames() ([]*domain.Shtemaran, domain.Error) {
	return p.db.GetShtemNames()
}

func NewShtemsRepository(db postgresrepository.ShtemsDB) *shtemsRepository {
	return &shtemsRepository{db}
}

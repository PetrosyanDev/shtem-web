package repositories

import (
	"shtem-web/sources/internal/core/domain"
	postgresrepository "shtem-web/sources/internal/repositories/postgres"
)

type shtemsRepository struct {
	db postgresrepository.ShtemsDB
}

func (p *shtemsRepository) GetShtems() ([]*domain.Shtemaran, domain.Error) {
	return p.db.GetShtems()
}
func (p *shtemsRepository) GetShtemLinkNames() ([]string, domain.Error) {
	return p.db.GetShtemLinkNames()
}
func (p *shtemsRepository) GetShtemByLinkName(name string) (*domain.Shtemaran, domain.Error) {
	return p.db.GetShtemByLinkName(name)
}

func (p *shtemsRepository) AllURLs() (*domain.SiteMapURLs, domain.Error) {
	return p.db.AllURLs()
}

func NewShtemsRepository(db postgresrepository.ShtemsDB) *shtemsRepository {
	return &shtemsRepository{db}
}

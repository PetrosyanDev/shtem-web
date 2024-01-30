package repositories

import (
	"shtem-web/sources/internal/core/domain"
	postgresrepository "shtem-web/sources/internal/repositories/postgres"
)

type categoriesRepository struct {
	db postgresrepository.CategoriesDB
}

func (p *categoriesRepository) GetCategories() ([]*domain.Category, domain.Error) {
	return p.db.GetCategories()
}

func NewCategoriesRepository(db postgresrepository.CategoriesDB) *categoriesRepository {
	return &categoriesRepository{db}
}

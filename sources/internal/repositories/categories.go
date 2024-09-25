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
func (p *categoriesRepository) GetCategoryByLinkName(c_link_name string) (*domain.Category, domain.Error) {
	return p.db.GetCategoryByLinkName(c_link_name)
}
func (p *categoriesRepository) GetCategoryByID(id int64) (*domain.Category, domain.Error) {
	return p.db.GetCategoryByID(id)
}

func NewCategoriesRepository(db postgresrepository.CategoriesDB) *categoriesRepository {
	return &categoriesRepository{db}
}

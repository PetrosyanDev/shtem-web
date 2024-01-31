package services

import (
	"shtem-web/sources/internal/core/domain"
	"shtem-web/sources/internal/repositories"
)

type categoriesService struct {
	categoriesRepository repositories.CategoriesRepository
}

func (q *categoriesService) GetCategories() ([]*domain.Category, domain.Error) {
	return q.categoriesRepository.GetCategories()
}
func (q *categoriesService) GetCategoriesWithShtems() (domain.Categories, domain.Error) {
	return q.categoriesRepository.GetCategoriesWithShtems()
}
func (q *categoriesService) GetShtemsByCategoryLinkName(c_linkName string) ([]*domain.Shtemaran, domain.Error) {
	return q.categoriesRepository.GetShtemsByCategoryLinkName(c_linkName)
}

func NewCategoriesService(categoriesRepository repositories.CategoriesRepository) *categoriesService {
	return &categoriesService{categoriesRepository}
}

package ports

import "shtem-web/sources/internal/core/domain"

type CategoriesService interface {
	GetCategories() ([]*domain.Category, domain.Error)
	GetCategoriesWithShtems() (domain.Categories, domain.Error)
}

// Erik Petrosyan ©
package postgresrepository

import "shtem-web/sources/internal/core/domain"

type QuestionsDB interface {
	GetShtemNames() ([]string, domain.Error)
}

type ShtemsDB interface {
	GetShtems() ([]*domain.Shtemaran, domain.Error)
	GetShtemsByCategoryId(c_id int64) ([]*domain.Shtemaran, domain.Error)
	GetShtemLinkNames() ([]string, domain.Error)
	GetShtemByLinkName(name string) (*domain.Shtemaran, domain.Error)
}

type CategoriesDB interface {
	GetCategories() ([]*domain.Category, domain.Error)
	GetCategoriesWithShtems() (domain.Categories, domain.Error)
}

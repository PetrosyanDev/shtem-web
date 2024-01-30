package repositories

import "shtem-web/sources/internal/core/domain"

type QuestionsRepository interface {
	GetShtemNames() ([]string, domain.Error)
}

type ShtemsRepository interface {
	GetShtems() ([]*domain.Shtemaran, domain.Error)
	GetShtemLinkNames() ([]string, domain.Error)
	GetShtemByLinkName(name string) (*domain.Shtemaran, domain.Error)
	AllURLs() (*domain.SiteMapURLs, domain.Error)
}

type CategoriesRepository interface {
	GetCategories() ([]*domain.Category, domain.Error)
}

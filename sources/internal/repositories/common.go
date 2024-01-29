package repositories

import "shtem-web/sources/internal/core/domain"

type QuestionsRepository interface {
	GetShtemNames() ([]string, domain.Error)
}

type ShtemsRepository interface {
	GetShtemNames() ([]*domain.Shtemaran, domain.Error)
	GetShtemByLinkName(name string) (*domain.Shtemaran, domain.Error)
	AllURLs() (*domain.SiteMapURLs, domain.Error)
}

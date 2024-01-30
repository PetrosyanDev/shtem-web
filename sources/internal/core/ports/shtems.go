package ports

import "shtem-web/sources/internal/core/domain"

type ShtemsService interface {
	GetShtems() ([]*domain.Shtemaran, domain.Error)
	GetShtemLinkNames() ([]string, domain.Error)
	GetShtemByLinkName(name string) (*domain.Shtemaran, domain.Error)
	GetSiteMap() ([]byte, domain.Error)
}

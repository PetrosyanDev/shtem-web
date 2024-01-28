package ports

import "shtem-web/sources/internal/core/domain"

type ShtemsService interface {
	GetShtemNames() ([]*domain.Shtemaran, domain.Error)
	GetShtemByLinkName(name string) (*domain.Shtemaran, domain.Error)
}

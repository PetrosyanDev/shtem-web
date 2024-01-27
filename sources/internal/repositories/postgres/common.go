// Erik Petrosyan ©
package postgresrepository

import "shtem-web/sources/internal/core/domain"

type QuestionsDB interface {
	GetShtemNames() ([]string, domain.Error)
}

type ShtemsDB interface {
	GetShtemNames() ([]*domain.Shtemaran, domain.Error)
}

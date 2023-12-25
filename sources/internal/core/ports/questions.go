package ports

import "shtem-web/sources/internal/core/domain"

type QuestionsService interface {
	GetShtemNames() ([]string, domain.Error)
}

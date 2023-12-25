package repositories

import "shtem-web/sources/internal/core/domain"

type QuestionsRepository interface {
	GetShtemNames() ([]string, domain.Error)
}
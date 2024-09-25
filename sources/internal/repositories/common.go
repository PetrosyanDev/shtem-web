package repositories

import "shtem-web/sources/internal/core/domain"

type QuestionsRepository interface {
	Create(question *domain.Question) domain.Error
	Update(question *domain.Question) domain.Error
	Delete(id int64) domain.Error
	FindBajin(question *domain.Question) ([]*domain.Question, domain.Error)
	FindQuestion(question *domain.Question) (*domain.Question, domain.Error)
	FindByID(id int) (*domain.Question, domain.Error)
}

type ShtemsRepository interface {
	GetShtems() ([]*domain.Shtemaran, domain.Error)
	GetShtemsByCategoryId(c_id int64) ([]*domain.Shtemaran, domain.Error)
	GetShtemLinkNames() ([]string, domain.Error)
	GetShtemByLinkName(name string) (*domain.Shtemaran, domain.Error)
}

type CategoriesRepository interface {
	GetCategories() ([]*domain.Category, domain.Error)
	GetCategoryByLinkName(linkName string) (*domain.Category, domain.Error)
	GetCategoryByID(id int64) (*domain.Category, domain.Error)
}

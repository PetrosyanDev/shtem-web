// Erik Petrosyan ©
package postgresrepository

import "shtem-web/sources/internal/core/domain"

type QuestionsDB interface {
	Create(question *domain.Question) domain.Error
	Update(question *domain.Question) domain.Error
	Delete(id int64) domain.Error
	FindQuestion(question *domain.Question) (*domain.Question, domain.Error)
	FindBajin(question *domain.Question) ([]*domain.Question, domain.Error)
	FindByID(id int) (*domain.Question, domain.Error)
}

type ShtemsDB interface {
	GetShtems() ([]*domain.Shtemaran, domain.Error)
	GetShtemsByCategoryId(c_id int64) ([]*domain.Shtemaran, domain.Error)
	GetShtemLinkNames() ([]string, domain.Error)
	GetShtemByLinkName(name string) (*domain.Shtemaran, domain.Error)
}

type CategoriesDB interface {
	GetCategories() ([]*domain.Category, domain.Error)
	GetCategoryByLinkName(c_link_name string) (*domain.Category, domain.Error)
	GetCategoriesWithShtems() (domain.Categories, domain.Error)
	GetShtemsByCategoryLinkName(c_linkName string) ([]*domain.Shtemaran, domain.Error)
}

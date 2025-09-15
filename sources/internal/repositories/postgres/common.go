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
	GetCategoryByLinkName(linkName string) (*domain.Category, domain.Error)
	GetCategoryByID(id int64) (*domain.Category, domain.Error)
}

type EmailsDB interface {
	InsertEmail(email string) domain.Error
	GetAllEmails() ([]*domain.Email, domain.Error)
}

type SponsorHitsDB interface {
	InsertSponsorHit(path, sponsorURL, clientID, ipHash, ua string) domain.Error
	GetDailyUniqueCount(path string, date string) (int64, domain.Error)
	GetDistinctPaths() ([]string, domain.Error)
}

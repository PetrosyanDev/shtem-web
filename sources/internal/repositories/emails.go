package repositories

import (
	"shtem-web/sources/internal/core/domain"
	postgresrepository "shtem-web/sources/internal/repositories/postgres"
)

type emailsRepository struct {
	db postgresrepository.EmailsDB
}

func (p *emailsRepository) InsertEmail(email string) domain.Error {
	return p.db.InsertEmail(email)
}

func (p *emailsRepository) GetAllEmails() ([]*domain.Email, domain.Error) {
	return p.db.GetAllEmails()
}

func NewEmailsRepository(db postgresrepository.EmailsDB) *emailsRepository {
	return &emailsRepository{db}
}

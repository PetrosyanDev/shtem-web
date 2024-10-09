package services

import (
	"shtem-web/sources/internal/core/domain"
	"shtem-web/sources/internal/repositories"
)

type emailsService struct {
	emailsRepository repositories.EmailsRepository
}

func (q *emailsService) InsertEmail(email string) domain.Error {
	return q.emailsRepository.InsertEmail(email)
}

func (q *emailsService) GetAllEmails() ([]*domain.Email, domain.Error) {
	return q.emailsRepository.GetAllEmails()
}

func NewEmailsService(emailsRepository repositories.EmailsRepository) *emailsService {
	return &emailsService{emailsRepository}
}

package ports

import "shtem-web/sources/internal/core/domain"

// EmailsService defines the methods for the emails service
type EmailsService interface {
	InsertEmail(email string) domain.Error
	GetAllEmails() ([]*domain.Email, domain.Error)
}

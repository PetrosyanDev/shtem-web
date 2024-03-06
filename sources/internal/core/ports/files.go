package ports

import (
	"shtem-web/sources/internal/core/domain"
)

type FilesService interface {
	Download(file *domain.File) domain.Error
}

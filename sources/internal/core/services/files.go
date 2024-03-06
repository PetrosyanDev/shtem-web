package services

import (
	storageclient "shtem-web/sources/internal/clients/storage"
	"shtem-web/sources/internal/core/domain"
)

type filesService struct {
	storage *storageclient.StorageClient
}

func (f *filesService) Download(file *domain.File) domain.Error {
	return f.storage.ProvideFile(file)
}

func NewFilesService(storage *storageclient.StorageClient) *filesService {
	return &filesService{storage}
}

package services

import (
	"shtem-web/sources/internal/core/domain"
	"shtem-web/sources/internal/repositories"
)

type sponsorHitsService struct {
	sponsorHitsRepository repositories.SponsorHitsRepository
}

func (s *sponsorHitsService) InsertSponsorHit(path, sponsorURL, clientID, ipHash, ua string) domain.Error {
	return s.sponsorHitsRepository.InsertSponsorHit(path, sponsorURL, clientID, ipHash, ua)
}

func (s *sponsorHitsService) GetDailyUniqueCount(path string, date string) (int64, domain.Error) {
	return s.sponsorHitsRepository.GetDailyUniqueCount(path, date)
}

func (s *sponsorHitsService) GetDistinctPaths() ([]string, domain.Error) {
	return s.sponsorHitsRepository.GetDistinctPaths()
}

func NewSponsorHitsService(repo repositories.SponsorHitsRepository) *sponsorHitsService {
	return &sponsorHitsService{repo}
}

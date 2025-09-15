package repositories

import (
	"shtem-web/sources/internal/core/domain"
	postgresrepository "shtem-web/sources/internal/repositories/postgres"
)

type sponsorHitsRepository struct {
	db postgresrepository.SponsorHitsDB
}

func (r *sponsorHitsRepository) InsertSponsorHit(path, sponsorURL, clientID, ipHash, ua string) domain.Error {
	return r.db.InsertSponsorHit(path, sponsorURL, clientID, ipHash, ua)
}

func (r *sponsorHitsRepository) GetDailyUniqueCount(path string, date string) (int64, domain.Error) {
	return r.db.GetDailyUniqueCount(path, date)
}

func (r *sponsorHitsRepository) GetDistinctPaths() ([]string, domain.Error) {
	return r.db.GetDistinctPaths()
}

func NewSponsorHitsRepository(db postgresrepository.SponsorHitsDB) *sponsorHitsRepository {
	return &sponsorHitsRepository{db}
}

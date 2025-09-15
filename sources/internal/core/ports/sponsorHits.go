package ports

import "shtem-web/sources/internal/core/domain"

type SponsorHitsService interface {
	InsertSponsorHit(path, sponsorURL, clientID, ipHash, ua string) domain.Error
	GetDailyUniqueCount(path string, date string) (int64, domain.Error)
	GetDistinctPaths() ([]string, domain.Error)
}

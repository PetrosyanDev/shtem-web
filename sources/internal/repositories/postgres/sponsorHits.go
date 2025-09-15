// Erik Petrosyan ©
package postgresrepository

import (
	"context"
	"database/sql"

	postgresclient "shtem-web/sources/internal/clients/postgres"
	"shtem-web/sources/internal/core/domain"
)

type sponsorHitsDB struct {
	ctx context.Context
	db  *postgresclient.PostgresDB
}

func NewSponsorHitsDB(ctx context.Context, db *postgresclient.PostgresDB) *sponsorHitsDB {
	return &sponsorHitsDB{ctx, db}
}

func (q *sponsorHitsDB) InsertSponsorHit(path, sponsorURL, clientID, ipHash, ua string) domain.Error {
	_, err := q.db.Exec(q.ctx, `
		INSERT INTO sponsor_hits (path, sponsor_url, client_id, ip_hash, user_agent, hit_date)
		VALUES ($1, $2, $3, $4, $5, CURRENT_DATE)
		ON CONFLICT ON CONSTRAINT uniq_daily_sponsor DO NOTHING
	`, path, sponsorURL, clientID, ipHash, ua)

	if err != nil {
		return domain.NewError().SetError(err)
	}
	return nil
}

// Optional: get daily unique users per path
func (q *sponsorHitsDB) GetDailyUniqueCount(path string, date string) (int64, domain.Error) {
	var count sql.NullInt64
	err := q.db.QueryRow(q.ctx, `
		SELECT COUNT(DISTINCT client_id)
		FROM sponsor_hits
		WHERE path = $1 AND hit_date = $2
	`, path, date).Scan(&count)

	if err != nil {
		return 0, domain.NewError().SetError(err)
	}
	return count.Int64, nil
}

func (q *sponsorHitsDB) GetDistinctPaths() ([]string, domain.Error) {
	rows, err := q.db.Query(q.ctx, `
		SELECT DISTINCT path FROM sponsor_hits
	`)
	if err != nil {
		return nil, domain.NewError().SetError(err)
	}
	defer rows.Close()

	var paths []string
	for rows.Next() {
		var p string
		if err := rows.Scan(&p); err != nil {
			return nil, domain.NewError().SetError(err)
		}
		paths = append(paths, p)
	}
	return paths, nil
}

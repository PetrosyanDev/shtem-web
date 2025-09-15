package reporting

import (
	"fmt"
	"log"
	"strings"
	"time"

	telegramclient "shtem-web/sources/internal/clients/telegram"
	"shtem-web/sources/internal/core/ports"
)

const (
	hour = 13
	min  = 0
	sec  = 0
)

type SponsorReporter struct {
	svc            ports.SponsorHitsService
	telegramClient *telegramclient.Client
}

func NewSponsorReporter(svc ports.SponsorHitsService, tg *telegramclient.Client) *SponsorReporter {
	return &SponsorReporter{svc, tg}
}

func (r *SponsorReporter) StartDaily() {
	for {
		now := time.Now()
		next := time.Date(now.Year(), now.Month(), now.Day(), hour, min, sec, 0, now.Location())
		if now.After(next) {
			next = next.Add(24 * time.Hour)
		}
		time.Sleep(time.Until(next))
		r.SendDailyReport()
	}
}

func (r *SponsorReporter) SendDailyReport() {
	date := time.Now().Format("2006-01-02")

	paths, err := r.svc.GetDistinctPaths()
	if err != nil {
		log.Println("failed to get paths:", err)
		return
	}

	var lines []string
	lines = append(lines, "📊 Daily Sponsor Report ("+date+")")

	for _, p := range paths {
		count, err := r.svc.GetDailyUniqueCount(p, date)
		if err != nil {
			log.Println("count error for", p, ":", err)
			continue
		}
		lines = append(lines, fmt.Sprintf("> %s = %d unique users", p, count))
	}

	msg := strings.Join(lines, "\n")
	r.telegramClient.Notify(msg)
}

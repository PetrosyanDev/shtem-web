package services

import (
	"bytes"
	"encoding/xml"
	"shtem-web/sources/internal/core/domain"
	"shtem-web/sources/internal/repositories"
	"time"
)

type shtemsService struct {
	shtemsRepository repositories.ShtemsRepository
}

func (q *shtemsService) GetShtems() ([]*domain.Shtemaran, domain.Error) {
	return q.shtemsRepository.GetShtems()
}
func (q *shtemsService) GetShtemsByCategoryId(c_id int64) ([]*domain.Shtemaran, domain.Error) {
	return q.shtemsRepository.GetShtemsByCategoryId(c_id)
}
func (q *shtemsService) GetShtemLinkNames() ([]string, domain.Error) {
	return q.shtemsRepository.GetShtemLinkNames()
}
func (q *shtemsService) GetShtemByLinkName(name string) (*domain.Shtemaran, domain.Error) {
	return q.shtemsRepository.GetShtemByLinkName(name)
}
func (q *shtemsService) GetSiteMap() ([]byte, domain.Error) {

	siteMap := new(domain.SiteMapURLs)

	// MAIN COMPONENTS
	siteMapHome := domain.SiteMapURL{
		Loc:        domain.BaseUrl,
		ChangeFreq: domain.SiteMapFreqDaily,
		LastMod:    time.Now().UTC().Format("2006-01-02"),
		Priority:   domain.SiteMapPriorityHighest,
	}

	siteMapShtems := domain.SiteMapURL{
		Loc:        domain.ShtemsUrl,
		ChangeFreq: domain.SiteMapFreqWeekly,
		LastMod:    time.Now().UTC().Format("2006-01-02"),
		Priority:   domain.SiteMapPriorityHigh,
	}

	siteMap.URLs = append(siteMap.URLs, siteMapHome, siteMapShtems)

	// SHTEMARANS

	allSingleShtemsURLs, err := q.shtemsRepository.GetShtemLinkNames()
	if err != nil {
		return nil, err
	}

	for _, url := range allSingleShtemsURLs {
		siteMap.URLs = append(siteMap.URLs, domain.SiteMapURL{
			Loc:        domain.ShtemsUrl + url,
			ChangeFreq: domain.SiteMapFreqMonthly,
			LastMod:    time.Now().UTC().Format("2006-01-02"),
			Priority:   domain.SiteMapPriorityMedium,
		})
	}

	// SHTEMARAN QUIZES
	for _, url := range allSingleShtemsURLs {
		siteMap.URLs = append(siteMap.URLs, domain.SiteMapURL{
			Loc:        domain.ShtemsUrl + url + "/quiz",
			ChangeFreq: domain.SiteMapFreqMonthly,
			LastMod:    time.Now().UTC().Format("2006-01-02"),
			Priority:   domain.SiteMapPriorityMedium,
		})
	}

	siteMap.XMLNS = domain.SiteMapXMLNS

	b := bytes.NewBuffer([]byte{})
	defer b.Reset()
	b.WriteString(xml.Header)
	enc := xml.NewEncoder(b)
	defer enc.Close()
	e := enc.Encode(siteMap)
	if err != nil {
		return nil, domain.NewError().SetError(e)
	}
	return b.Bytes(), nil
}

func NewShtemsService(shtemsRepository repositories.ShtemsRepository) *shtemsService {
	return &shtemsService{shtemsRepository}
}

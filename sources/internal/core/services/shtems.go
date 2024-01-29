package services

import (
	"bytes"
	"encoding/xml"
	"shtem-web/sources/internal/core/domain"
	"shtem-web/sources/internal/repositories"
)

type shtemsService struct {
	shtemsRepository repositories.ShtemsRepository
}

func (q *shtemsService) GetShtemNames() ([]*domain.Shtemaran, domain.Error) {
	return q.shtemsRepository.GetShtemNames()
}
func (q *shtemsService) GetShtemByLinkName(name string) (*domain.Shtemaran, domain.Error) {
	return q.shtemsRepository.GetShtemByLinkName(name)
}
func (q *shtemsService) GetSiteMap() ([]byte, domain.Error) {

	siteMap, err := q.shtemsRepository.AllURLs()
	if err != nil {
		return nil, err
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

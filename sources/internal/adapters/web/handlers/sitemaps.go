package handlers

import (
	"bytes"
	"encoding/xml"
	"log"
	"net/http"
	"time"

	"shtem-web/sources/internal/adapters/web/dto"
	"shtem-web/sources/internal/core/domain"

	"github.com/gin-gonic/gin"
)

func (h *webHandler) SiteMapForAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		siteMapData, err := h.getSiteMap()
		if err != nil {
			log.Printf("webHandler:SiteMap (%v)", err.RawError())
			ctx.XML(http.StatusInternalServerError, nil)
			return
		}
		dto.WriteXMLResponse(ctx, "sitemap.xml", &siteMapData)
	}
}

func (h *webHandler) getSiteMap() ([]byte, domain.Error) {

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
	allSingleShtemsURLs, err := h.shtemsService.GetShtemLinkNames()
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

	// SHTEMARAN PDFs
	allSingleShtemsPDFURLs, err := h.shtemsService.GetShtems()
	if err != nil {
		return nil, err
	}

	// SHTEMARAN QUIZES
	for _, url := range allSingleShtemsPDFURLs {
		if url.HasQuiz {
			siteMap.URLs = append(siteMap.URLs, domain.SiteMapURL{
				Loc:        domain.ShtemsUrl + url.LinkName + "/quiz",
				ChangeFreq: domain.SiteMapFreqMonthly,
				LastMod:    time.Now().UTC().Format("2006-01-02"),
				Priority:   domain.SiteMapPriorityMedium,
			})
		}
	}

	// CATEGORIES
	categories, err := h.categoriesService.GetCategories()
	if err != nil {
		return nil, err
	}

	for _, c := range categories {
		siteMap.URLs = append(siteMap.URLs, domain.SiteMapURL{
			Loc:        domain.CategoryUrl + c.LinkName,
			ChangeFreq: domain.SiteMapFreqMonthly,
			LastMod:    time.Now().UTC().Format("2006-01-02"),
			Priority:   domain.SiteMapPriorityHigh,
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

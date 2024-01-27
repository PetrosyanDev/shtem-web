package domain

import (
	"shtem-web/sources/internal/utils"
	"slices"
	"strings"
)

type OpenGraphTAG struct {
	URL, Type, Title, Description, Image, Video string
}

type TwitterTAG struct {
	Card, Domain, URL, Title, Description, Image, Video string
}

type Header struct {
	Title        string
	Description  string
	Keywords     string
	AppCapable   string
	OpenGraphTAG OpenGraphTAG
	TwitterTAG   TwitterTAG
	IconLarge    string
	IconSmall    string
}

func (h *Header) PopulateKeywords(kwds ...string) {
	keywords := []string{}
	for _, v := range utils.GenKWDs(h.Title, false) {
		keywords = append(keywords, v)
	}
	for _, v := range utils.GenKWDs(h.Description, false) {
		keywords = append(keywords, v)
	}
	for _, v := range kwds {
		keywords = append(keywords, v)
	}
	h.Keywords = strings.Join(slices.Compact[[]string](keywords), ", ")
}

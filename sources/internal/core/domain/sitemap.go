package domain

import "encoding/xml"

const SiteMapXMLNS = "http://www.sitemaps.org/schemas/sitemap/0.9"

const (
	SiteMapFreqAlways  = "always"
	SiteMapFreqHourly  = "hourly"
	SiteMapFreqDaily   = "daily"
	SiteMapFreqWeekly  = "weekly"
	SiteMapFreqMonthly = "monthly"
	SiteMapFreqYearly  = "yearly"
	SiteMapFreqNever   = "never"
)

const (
	SiteMapPriorityHighest = "1.0"
	SiteMapPriorityHigh    = "0.9"
	SiteMapPriorityMedium  = "0.7"
	SiteMapPriorityLow     = "0.5"
	SiteMapPriorityLowest  = "0.2"
)

type SiteMapURL struct {
	XMLName    xml.Name `xml:"url"`
	Loc        string   `xml:"loc"`
	ChangeFreq string   `xml:"changefreq"`
	LastMod    string   `xml:"lastmod"`
	Priority   string   `xml:"priority"`
}

type SiteMapURLs struct {
	XMLName xml.Name     `xml:"urlset"`
	XMLNS   string       `xml:"xmlns,attr"`
	URLs    []SiteMapURL `xml:"url"`
}

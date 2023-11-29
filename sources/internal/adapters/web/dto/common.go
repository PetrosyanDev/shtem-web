// Erik Petrosyan ©
package dto

import (
	"fmt"
	"shtem-web/sources/internal/core/domain"
)

type pageBuilder struct {
	page        *domain.Page
	domain, url string
}

func FormatRelativeURL(path *string) string {
	return fmt.Sprintf("cdn/%s", *path)
}

func (b *pageBuilder) AddHeader(title, description string) *pageBuilder {
	b.page.Header.Title = title
	b.page.Header.Description = description
	return b
}

func (b *pageBuilder) AddTopMenuItem(name, link string, isCurrent bool) *pageBuilder {
	b.page.Body.TopMenu.Items = append(b.page.Body.TopMenu.Items, domain.TopMenuItem{
		Name: name, Link: link, IsCurrent: isCurrent,
	})
	if isCurrent {
		b.page.Body.TopMenu.CurrentName = name
	}
	return b
}

func (b *pageBuilder) Page() *domain.Page {
	return b.page
}

func newPageBuilder() *pageBuilder {
	d, url := headerDomain, headerURL
	return &pageBuilder{
		&domain.Page{}, d, url,
	}
}

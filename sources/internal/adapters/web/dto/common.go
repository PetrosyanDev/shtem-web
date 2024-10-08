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

func (b *pageBuilder) AddHeader(title, description, app string, keywords ...string) *pageBuilder {
	b.page.Header.Title = title
	b.page.Header.Description = description
	b.page.Header.AppCapable = app
	b.page.Header.IconLarge = headerIconLarge
	b.page.Header.IconSmall = headerIconSmall
	keywords = append(keywords, headerDefaultKwds...)
	b.page.Header.PopulateKeywords(keywords...)

	return b
}

func (b *pageBuilder) AddOpenGraphTAG(title, description, path, image string) *pageBuilder {
	b.page.Header.OpenGraphTAG.URL = fmt.Sprintf("%s%s", b.url, path)
	b.page.Header.OpenGraphTAG.Type = headerType
	b.page.Header.OpenGraphTAG.Title = title
	b.page.Header.OpenGraphTAG.Description = description
	if image != "" {
		b.page.Header.OpenGraphTAG.Image = fmt.Sprintf("%s/%s", b.url, image)
	}
	return b
}

func (b *pageBuilder) AddTwitterTAG(title, description, path, image string) *pageBuilder {
	b.page.Header.TwitterTAG.Card = headerTwitterCard
	b.page.Header.TwitterTAG.Domain = b.domain
	b.page.Header.TwitterTAG.URL = fmt.Sprintf("%s%s", b.url, path)
	b.page.Header.TwitterTAG.Title = title
	b.page.Header.TwitterTAG.Description = description
	if image != "" {
		b.page.Header.TwitterTAG.Image = fmt.Sprintf("%s/%s", b.url, image)
	}
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

func (b *pageBuilder) AddSingleShtem(shtemaran *domain.Shtemaran) *pageBuilder {

	b.page.Body.CurrentShtem = domain.Shtemaran{
		Name:        shtemaran.Name,
		Description: shtemaran.Description,
		Author:      shtemaran.Author,
		LinkName:    shtemaran.LinkName,
		Image:       shtemaran.Image,
		PDF:         shtemaran.PDF,
		HasQuiz:     shtemaran.HasQuiz,
		HasPDF:      shtemaran.HasPDF,
	}

	return b
}

func (b *pageBuilder) AddShtemNames(shtems []*domain.Shtemaran) *pageBuilder {
	for _, n := range shtems {
		b.page.Body.Shtems = append(b.page.Body.Shtems, domain.Shtemaran{
			Name:        n.Name,
			Description: n.Description,
			Author:      n.Author,
			LinkName:    n.LinkName,
			Image:       n.Image,
			PDF:         n.PDF,
			HasQuiz:     n.HasQuiz,
			HasPDF:      n.HasPDF,
		})
	}

	return b
}

// CATEGORIES

func (b *pageBuilder) AddCategories(categories *domain.Categories) *pageBuilder {

	allCategories := domain.CategoriesTpl{}

	for _, value := range *categories {

		shtemarans := []domain.Shtemaran{}
		for _, shtem := range value.Shtemarans {
			shtemarans = append(shtemarans, *shtem)
		}
		allCategories = append(allCategories, domain.SortedCategoryTpl{
			Category:   *value.Category,
			Shtemarans: shtemarans,
		})
	}
	b.page.Body.Categories = allCategories

	return b
}

func (b *pageBuilder) AddSingleCategory(category *domain.Category) *pageBuilder {

	b.page.Body.CurrentCategory = domain.Category{
		Name:        category.Name,
		Description: category.Description,
		LinkName:    category.LinkName,
	}

	return b
}

func (b *pageBuilder) Page() *domain.Page {
	return b.page
}

func (b *pageBuilder) AddCategoryShtemNames(shtems []*domain.Shtemaran) *pageBuilder {
	for _, n := range shtems {
		b.page.Body.CurrentCategoryShtems = append(b.page.Body.CurrentCategoryShtems, domain.Shtemaran{
			Name:        n.Name,
			Description: n.Description,
			Author:      n.Author,
			LinkName:    n.LinkName,
			Image:       n.Image,
			PDF:         n.PDF,
			HasQuiz:     n.HasQuiz,
			HasPDF:      n.HasPDF,
		})
	}

	return b
}

func newPageBuilder() *pageBuilder {
	d, url := headerDomain, headerURL
	return &pageBuilder{
		&domain.Page{}, d, url,
	}
}

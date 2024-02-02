// Erik Petrosyan ©
package dto

import (
	"shtem-web/sources/internal/core/domain"
)

func SingleShtemData(category *domain.Category, shtemaran *domain.Shtemaran) *domain.Page {
	var (
		title       = category.Name + " - " + shtemaran.Name + " • shtemaran.am"
		description = shtemaran.Description
		socImage    = headerDefaultSocialImage
	)
	const (
		app  = "no"
		path = ""
	)

	kwds := append([]string{}, shtemaran.Keywords...)
	pb := newPageBuilder().
		AddHeader(title, description, app, kwds...).
		AddOpenGraphTAG(title, description, path, socImage).
		AddTwitterTAG(title, description, path, socImage).
		AddTopMenuItem("ԳԼԽԱՎՈՐ", "/", false).
		AddTopMenuItem("ՇՏԵՄԵՐ", "/shtems", true).
		AddSingleShtem(shtemaran)

	return pb.Page()

}

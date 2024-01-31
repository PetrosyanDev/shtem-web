package dto

import "shtem-web/sources/internal/core/domain"

func CategoryData(category *domain.Category, shtemarans []*domain.Shtemaran) *domain.Page {
	var (
		title       = category.Name + " • shtemaran.am"
		description = category.Description
		socImage    = headerDefaultSocialImage
	)
	const (
		app  = "no"
		path = ""
	)

	kwds := []string{}
	pb := newPageBuilder().
		AddHeader(title, description, app, kwds...).
		AddOpenGraphTAG(title, description, path, socImage).
		AddTwitterTAG(title, description, path, socImage).
		AddTopMenuItem("ԳԼԽԱՎՈՐ", "/", false).
		AddTopMenuItem("ՇՏԵՄԵՐ", "/shtems", false).
		AddSingleCategory(category).
		AddCategoryShtemNames(shtemarans)

	return pb.Page()

}

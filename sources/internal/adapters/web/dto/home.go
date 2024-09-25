// Erik Petrosyan ©
package dto

import "shtem-web/sources/internal/core/domain"

func HomeData(shtemNames []*domain.Shtemaran, categories *domain.Categories) *domain.Page {
	const (
		title       = "shtemaran.am • Սովորիր արագ | Ժամանակը խնայելու լավագույն միջոցը"
		description = "Բարի գալուստ shtemaran.am, այստեղ դուք կգտնեք ձեր նախընտրած ցանկացած շտեմարան"
		app         = "no"
		path        = ""
		socImage    = headerDefaultSocialImage
	)

	kwds := []string{}
	pb := newPageBuilder().
		AddHeader(title, description, app, kwds...).
		AddOpenGraphTAG(title, description, path, socImage).
		AddTwitterTAG(title, description, path, socImage).
		AddTopMenuItem("ԳԼԽԱՎՈՐ", "/", true).
		// AddTopMenuItem("ՇՏԵՄԵՐ", "/shtems", false).
		AddShtemNames(shtemNames).
		AddCategories(categories)

	return pb.Page()

}

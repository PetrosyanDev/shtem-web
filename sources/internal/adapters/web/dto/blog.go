// Erik Petrosyan ©
package dto

import "shtem-web/sources/internal/core/domain"

func BlogData() *domain.Page {
	const (
		title       = "shtemaran.am Բլոգ • Սովորիր արագ | Ժամանակը խնայելու լավագույն միջոցը"
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
		AddTopMenuItem("ԳԼԽԱՎՈՐ", "/", false).
		AddTopMenuItem("ԲԼՈԳ", "/blog", true)
	return pb.Page()
}

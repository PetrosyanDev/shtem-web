// Erik Petrosyan ©
package dto

import (
	"shtem-web/sources/internal/core/domain"
)

type ShtemsResponse struct {
	Data []string `json:"data"`
}

func ShtemsData(shtemNames []*domain.Shtemaran) *domain.Page {
	const (
		title       = "shtemaran.am • Learning Fast | The Best Way to Save Time"
		description = "Welcome to shtemaran.am"
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
		AddTopMenuItem("ՇՏԵՄԵՐ", "/shtems", true).
		AddShtemNames(shtemNames)

	return pb.Page()

}

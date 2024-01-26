// Erik Petrosyan ©
package dto

import (
	"shtem-web/sources/internal/core/domain"
)

type ShtemsResponse struct {
	Data []string `json:"data"`
}

func ShtemsData(shtemNames []string) *domain.Page {
	const (
		title       = "shtemaran.am • Learning Fast | The Best Way to Save Time"
		description = "Welcome to shtemaran.am"
	)

	pb := newPageBuilder().
		AddHeader(title, description).
		AddTopMenuItem("ԳԼԽԱՎՈՐ", "/", false).
		AddTopMenuItem("ՇՏԵՄԵՐ", "/shtems", true).
		AddTopMenuItem("ԻՄ ՄԱՍԻՆ", "/about", false).
		AddShtemNames(shtemNames)

	return pb.Page()

}

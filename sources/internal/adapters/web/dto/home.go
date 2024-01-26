// Erik Petrosyan ©
package dto

import "shtem-web/sources/internal/core/domain"

func HomeData() *domain.Page {
	const (
		title       = "shtemaran.am • Learning Fast | The Best Way to Save Time"
		description = "Welcome to shtemaran.am"
	)
	pb := newPageBuilder().
		AddHeader(title, description).
		AddTopMenuItem("ԳԼԽԱՎՈՐ", "/", true).
		AddTopMenuItem("ՇՏԵՄԵՐ", "/shtems", false).
		AddTopMenuItem("ԻՄ ՄԱՍԻՆ", "/about", false)
	return pb.Page()
}

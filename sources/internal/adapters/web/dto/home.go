// Erik Petrosyan ©
package dto

import "shtem-web/sources/internal/core/domain"

func HomeData() *domain.Page {
	const (
		title       = "Shtemaran.am - Learing Fast | The best way to save up time"
		description = "Welcome to shtemaran.am"
	)
	pb := newPageBuilder().
		AddHeader(title, description)
	return pb.Page()
}

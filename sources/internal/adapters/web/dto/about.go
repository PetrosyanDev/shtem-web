// Erik Petrosyan ©
package dto

import "shtem-web/sources/internal/core/domain"

func AboutData() *domain.Page {
	const (
		title       = "shtemaran.am • Learning Fast | The Best Way to Save Time"
		description = "Welcome to shtemaran.am"
	)
	pb := newPageBuilder().
		AddHeader(title, description).
		AddTopMenuItem("HOME", "/", false).
		AddTopMenuItem("SHTEMS", "/shtems", false).
		AddTopMenuItem("ABOUT", "/about", true)
	return pb.Page()
}

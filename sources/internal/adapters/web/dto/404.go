// Erik Petrosyan ©
package dto

import "shtem-web/sources/internal/core/domain"

func NotFoundData() *domain.Page {
	const (
		title       = "Oops:( Something went wrong."
		description = "Error 404 - The resource you are locking does not exists or has been removed."
		app         = "no"
		path        = ""
		socImage    = headerDefaultSocialImage
	)

	pb := newPageBuilder().
		AddHeader(title, description, app).
		AddTopMenuItem("ԳԼԽԱՎՈՐ", "/", false)
	return pb.Page()
}

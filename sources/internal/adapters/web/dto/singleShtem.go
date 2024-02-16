// Erik Petrosyan ©
package dto

import (
	"shtem-web/sources/internal/core/domain"
	"strings"

	"golang.org/x/net/html"
)

func extractText(node *html.Node) string {
	var result string
	if node.Type == html.TextNode {
		result += node.Data
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		result += extractText(child)
	}
	return result
}

func convertDesc(description string) string {
	// description
	reader := strings.NewReader(description)
	doc, err := html.Parse(reader)
	if err != nil {
		return "shtemaran.am"
	}

	description = extractText(doc)

	if len(description) > 320 {
		description = description[:320]
	}

	return description
}

func SingleShtemData(category *domain.Category, shtemaran *domain.Shtemaran) *domain.Page {
	var (
		title       = category.Name + " - " + shtemaran.Name + " • shtemaran.am"
		description = convertDesc(shtemaran.Description)
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

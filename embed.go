// Erik Petrosyan ©
package embd

import (
	"embed"
)

var (
	//go:embed all:assets
	assets embed.FS
	//go:embed all:uploads
	uploads embed.FS
	//go:embed all:templates
	templates embed.FS
)

type _EMBD struct {
	Assets    *embed.FS
	Uploads   *embed.FS
	Templates *embed.FS
}

func NewEMBD() *_EMBD {
	return &_EMBD{
		Assets:    &assets,
		Uploads:   &uploads,
		Templates: &templates,
	}
}

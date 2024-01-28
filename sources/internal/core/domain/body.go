package domain

type TopMenuItem struct {
	Name      string
	Link      string
	IsCurrent bool
}

type TopMenu struct {
	CurrentName string
	Items       []TopMenuItem
}

type Body struct {
	TopMenu      TopMenu
	Shtems       []Shtemaran
	CurrentShtem Shtemaran
}

type Shtemaran struct {
	Name        string
	Description string
	Author      string
	LinkName    string
	Image       string
	PDF         string
}

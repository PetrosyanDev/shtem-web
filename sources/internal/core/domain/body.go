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
	TopMenu TopMenu
	Shtems  []Shtemaran
}

type Shtemaran struct {
	Name        string
	Description string
	LinkName    string
}

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

type Shtem struct {
	Name string
}

type Body struct {
	TopMenu TopMenu
	Shtems  []Shtem
}

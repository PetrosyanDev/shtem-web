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

type Categories map[Category][]*Shtemaran
type CategoriesTpl map[Category][]Shtemaran

type Body struct {
	TopMenu               TopMenu
	Shtems                []Shtemaran
	Categories            CategoriesTpl
	CurrentShtem          Shtemaran
	CurrentCategory       Category
	CurrentCategoryShtems []Shtemaran
}

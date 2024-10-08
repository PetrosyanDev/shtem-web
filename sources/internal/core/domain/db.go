package domain

import "time"

type Question struct {
	Q_id     int64
	Bajin    int
	Mas      int
	Q_number int
	Text     string
	Options  []string
	Answers  []int
	ShtemId  int64
}

type Shtemaran struct {
	Id          int64
	Name        string
	Description string
	Author      string
	LinkName    string
	Image       string
	PDF         string
	Keywords    []string
	Category    int64
	HasQuiz     bool
	HasPDF      bool
}

type Category struct {
	C_id        int64
	Name        string
	Description string
	LinkName    string
	Score       int
}

type SortedCategory struct {
	Category   *Category
	Shtemarans []*Shtemaran
}

type SortedCategoryTpl struct {
	Category   Category
	Shtemarans []Shtemaran
}

type Categories []SortedCategory
type CategoriesTpl []SortedCategoryTpl

type Email struct {
	Id        int64
	Email     string
	CreatedAt time.Time
}

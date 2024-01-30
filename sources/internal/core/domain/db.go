package domain

type Shtemaran struct {
	Id          int64
	Name        string
	Description string
	Author      string
	LinkName    string
	Image       string
	PDF         string
}

type Category struct {
	C_id        int64
	Name        string
	Description string
}

package models

type Meal struct {
	Id   int `pk:"auto;column(id)"`
	Cnt  int
	Loc  string
	Name string
}

type Groups struct {
	Id   int `pk:"auto;column(id)"`
	Date string
	Name string
}

type Record struct {
	Id      int `pk:"auto;column(id)"`
	Name    string
	Date    string
	Main    string
	Drink   string
	Groupid int
}

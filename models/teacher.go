package models

type Teacher struct {
	Id     int64
	Number string
	Name   string
	Sex    string
	Phone  string
	Qq     string
	CCT    []*CCT `orm:"reverse(many)"`
}

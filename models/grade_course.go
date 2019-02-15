package models

type GradeCourse struct {
	Id     int64
	Grade  *Grade  `orm:"rel(fk)"`
	Course *Course `orm:"rel(fk)"`
}

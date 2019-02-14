package models

type GradeCourse struct {
	Id       int64
	GradeId  *Grade  `orm:"rel(fk)"`
	CourseId *Course `orm:"rel(fk)"`
}

package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

func Init() {
	orm.RegisterModel(new(Course), new(Clazz), new(ClazzCourseTeacher), new(Escore), new(Exam),
		new(Grade), new(GradeCourse), new(Student), new(System), new(Teacher), new(User))
}

type Course struct {
	Id   int64
	Name string
}

type Clazz struct {
	Id    int64
	Name  string
	Grade *Grade `orm:"rel(fk)"`
}

type ClazzCourseTeacher struct {
	Id        int64
	ClazzId   *Clazz   `orm:"rel(fk)"`
	GradeId   *Grade   `orm:"rel(fk)"`
	CourseId  *Course  `orm:"rel(fk)"`
	TeacherId *Teacher `orm:"rel(fk)"`
}

type Escore struct {
	Id        int64
	ExamId    *Exam    `orm:"rel(fk)"`
	GradeId   *Grade   `orm:"rel(fk)"`
	CourseId  *Course  `orm:"rel(fk)"`
	ClazzId   *Clazz   `orm:"rel(fk)"`
	StudentId *Student `orm:"rel(fk)"`
	Score     int64
}

type Exam struct {
	Id       int64
	Name     string
	Time     time.Time
	Remark   string
	Type     int
	GradeId  *Grade  `orm:"rel(fk)"`
	ClazzId  *Clazz  `orm:"rel(fk)"`
	CourseId *Course `orm:"rel(fk)"`
}

type Grade struct {
	Id    int64
	Name  string
	Clazz []*Clazz `orm:"reverse(many)"`
}

type GradeCourse struct {
	Id       int64
	GradeId  *Grade  `orm:"rel(fk)"`
	CourseId *Course `orm:"rel(fk)"`
}

type Student struct {
	Id      int64
	Number  string
	Name    string
	Sex     string
	Phone   string
	Qq      string
	Photo   string
	ClazzId *Clazz `orm:"rel(fk)"`
	GradeId *Grade `orm:"rel(fk)"`
}

type System struct {
	Id            int64
	SchoolName    string
	ForbidTeacher int64
	ForbidStudent int64
	NoticeTeacher string
	NoticeStudent string
}

type Teacher struct {
	Id     int64
	Number string
	Name   string
	Sex    string
	Phone  string
	Qq     string
	Photo  string
}

type User struct {
	Id       int64
	Account  string
	Password string
	Name     string
	Type     int
}

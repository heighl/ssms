package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Exam struct {
	Id       int64
	Name     string
	Time     time.Time
	Remark   string
	Type     int64
	GradeId  *Grade    `orm:"rel(fk)"`
	ClazzId  *Clazz    `orm:"rel(fk)"`
	CourseId *Course   `orm:"rel(fk)"`
	Escore   []*Escore `orm:"reverse(many)"`
}

func (exam *Exam) OneGrade() {
	o := orm.NewOrm()
	err := o.Read(exam)
	if err != nil {
		return
	}
}

func (exam *Exam) AllGrade() ([]*Exam, error) {
	var exams []*Exam
	o := orm.NewOrm()
	_, err := o.QueryTable(Clazz{}).OrderBy("Id").All(&exams)
	if err != nil {
		return nil, err
	}
	return exams, nil
}

func (exam *Exam) Add() error {
	o := orm.NewOrm()
	if _, err := o.Insert(exam); err != nil {
		return err
	}
	return nil
}

func (exam *Exam) Update() error {
	o := orm.NewOrm()
	if _, err := o.Update(exam, "Name", "Time", "Remark", "Type", "GradeId", "ClazzId", "CourseId"); err != nil {
		return err
	}
	return nil
}

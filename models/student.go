package models

import "github.com/astaxie/beego/orm"

type Student struct {
	Id      int64
	Number  string
	Name    string
	Sex     string
	Phone   string
	Qq      string
	ClazzId *Clazz    `orm:"rel(fk)"`
	GradeId *Grade    `orm:"rel(fk)"`
	Escore  []*Escore `orm:"reverse(many)"`
}

func (student *Student) OneGrade() {
	o := orm.NewOrm()
	err := o.Read(student)
	if err != nil {
		return
	}
}

func (student *Student) AllGrade() ([]*Student, error) {
	var students []*Student
	o := orm.NewOrm()
	_, err := o.QueryTable(Student{}).OrderBy("Id").All(&student)
	if err != nil {
		return nil, err
	}
	return students, nil
}

func (student *Student) Add() error {
	o := orm.NewOrm()
	if _, err := o.Insert(student); err != nil {
		return err
	}
	return nil
}

func (student *Student) Update() error {
	o := orm.NewOrm()
	if _, err := o.Update(student, "Name"); err != nil {
		return err
	}
	return nil
}

package models

import (
	"github.com/apex/log"
	"github.com/astaxie/beego/orm"
)

type Student struct {
	Id      int
	Number  string
	Name    string
	Sex     string
	Phone   string
	Qq      string
	Clazz *Clazz    `orm:"rel(fk)"`
	Grade *Grade    `orm:"rel(fk)"`
	//Escore  []*Escore `orm:"reverse(many)"`
}

func (this *Student) GetOne() *Student {
	o := orm.NewOrm()
	err := o.Read(this)
	if err != nil {
		log.Info(err.Error())
		return nil
	}
	return this
}

func (student *Student) AllGrade() ([]*Student, error) {
	var students []*Student
	o := orm.NewOrm()
	_, err := o.QueryTable(Student{}).OrderBy("Id").All(&students)
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

package models

import "github.com/astaxie/beego/orm"

type Clazz struct {
	Id      int64
	Name    string
	GradeId *Grade     `orm:"rel(fk)"`
	CCT     []*CCT     `orm:"reverse(many)"`
	Escore  []*Escore  `orm:"reverse(many)"`
	Exam    []*Exam    `orm:"reverse(many)"`
	Student []*Student `orm:"reverse(many)"`
}

func (grade *Clazz) OneGrade() {
	o := orm.NewOrm()
	err := o.Read(grade)
	if err != nil {
		return
	}
}

func (grade *Clazz) AllGrade() ([]*Grade, error) {
	var grades []*Grade
	o := orm.NewOrm()
	_, err := o.QueryTable(Clazz{}).OrderBy("Id").All(&grades)
	if err != nil {
		return nil, err
	}
	return grades, nil
}

func (grade *Clazz) Add() error {
	o := orm.NewOrm()
	if _, err := o.Insert(grade); err != nil {
		return err
	}
	return nil
}

func (grade *Clazz) Update() error {
	o := orm.NewOrm()
	if _, err := o.Update(grade, "Name", "GradeId"); err != nil {
		return err
	}
	return nil
}

package models

import (
	"github.com/astaxie/beego/orm"
)

type Clazz struct {
	Id      int64
	Name    string
	GradeId *Grade `orm:"rel(fk)"`
	//CCT     []*CCT     `orm:"reverse(many)"`
	//Escore  []*Escore  `orm:"reverse(many)"`
	//Exam    []*Exam    `orm:"reverse(many)"`
	//Student []*Student `orm:"reverse(many)"`
}

func (clazz *Clazz) Grade() *Clazz {
	o := orm.NewOrm()
	err := o.Read(clazz)
	if err != nil {
		return nil
	}
	return clazz
}

func (clazz *Clazz) AllGrade() ([]*Clazz, error) {
	var clazzs []*Clazz
	o := orm.NewOrm()
	_, err := o.QueryTable("clazz").OrderBy("Id").All(&clazzs)
	if err != nil {
		return nil, err
	}
	return clazzs, nil
}

func (clazz *Clazz) Add() error {
	o := orm.NewOrm()
	if _, err := o.Insert(clazz); err != nil {
		return err
	}
	return nil
}

func (clazz *Clazz) Update() error {
	o := orm.NewOrm()
	if _, err := o.Update(clazz, "Name", "GradeId"); err != nil {
		return err
	}
	return nil
}

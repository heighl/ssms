package models

import (
	"github.com/apex/log"
	"github.com/astaxie/beego/orm"
)

type Teacher struct {
	Id     int64
	Number string
	Name   string
	Sex    string
	Phone  string
	Qq     string
	//CCT    []*CCT `orm:"reverse(many)"`
}

func (this *Teacher) GetOne() *Teacher {
	o := orm.NewOrm()
	err := o.Read(this)
	if err != nil {
		log.Info(err.Error())
		return nil
	}
	return this
}

func (this *Teacher) AllGrade() ([]*Teacher, error) {
	var teachers []*Teacher
	o := orm.NewOrm()
	_, err := o.QueryTable(Teacher{}).OrderBy("Id").All(&teachers)
	if err != nil {
		return nil, err
	}
	return teachers, nil
}

func (this *Teacher) Add() error {
	o := orm.NewOrm()
	if _, err := o.Insert(this); err != nil {
		return err
	}
	return nil
}

func (this *Teacher) Update() error {
	o := orm.NewOrm()
	if _, err := o.Update(this, "Name"); err != nil {
		return err
	}
	return nil
}

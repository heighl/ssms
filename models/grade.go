package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Grade struct {
	Id          int64
	Name        string
	//Clazz       []*Clazz       `orm:"reverse(many)"`
	//CCT         []*CCT         `orm:"reverse(many)"`
	//Escore      []*Escore      `orm:"reverse(many)"`
	Exam        []*Exam        `orm:"reverse(many)"`
	//GradeCourse []*GradeCourse `orm:"reverse(many)"`
	//Student     []*Student     `orm:"reverse(many)"`
}

func (grade *Grade) OneGrade() *Grade {
	o := orm.NewOrm()
	err := o.Read(grade)
	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
		return grade
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
		return grade
	} else {
		fmt.Println(grade.Id, grade.Name)
		return grade
	}
}

func (grade *Grade) AllGrade() ([]*Grade, error) {
	var grades []*Grade
	o := orm.NewOrm()
	_, err := o.QueryTable(Grade{}).OrderBy("Id").All(&grades)
	if err != nil {
		return nil, err
	}
	return grades, nil
}

func (grade *Grade) Add() error {
	o := orm.NewOrm()
	if _, err := o.Insert(grade); err != nil {
		return err
	}
	return nil
}

func (grade *Grade) Update() error {
	o := orm.NewOrm()
	if _, err := o.Update(grade, "Name"); err != nil {
		return err
	}
	return nil
}

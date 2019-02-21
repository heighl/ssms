package models

import "github.com/astaxie/beego/orm"

type Course struct {
	Id          int64
	Name        string
	//CCT         []*CCT         `orm:"reverse(many)"`
	//Escore      []*Escore      `orm:"reverse(many)"`
	Exam        []*Exam        `orm:"reverse(many)"`
	//GradeCourse []*GradeCourse `orm:"reverse(many)"`
}

func (course *Course) OneCourse() *Course {
	o := orm.NewOrm()
	err := o.Read(course)
	if err != nil {
		return nil
	}
	return course
}

func (course *Course) AllCourse() ([]*Course, error) {
	var courses []*Course
	o := orm.NewOrm()
	_, err := o.QueryTable(Course{}).OrderBy("Id").RelatedSel().All(&courses)
	if err != nil {
		return nil, err
	}
	return courses, nil
}

func (course *Course) Add() error {
	o := orm.NewOrm()
	if _, err := o.Insert(course); err != nil {
		return err
	}
	return nil
}

func (course *Course) Update() error {
	o := orm.NewOrm()
	if _, err := o.Update(course, "Name"); err != nil {
		return err
	}
	return nil
}

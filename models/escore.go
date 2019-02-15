package models

import "github.com/astaxie/beego/orm"

type Escore struct {
	Id      int64
	Exam    *Exam    `orm:"rel(fk)"`
	Clazz   *Clazz   `orm:"rel(fk)"`
	Student *Student `orm:"rel(fk)"`
	Grade   *Grade   `orm:"rel(fk)"`
	Course  *Course  `orm:"rel(fk)"`
	Score   int64
}

func (escore *Escore) OneGrade() {
	o := orm.NewOrm()
	err := o.Read(escore)
	if err != nil {
		return
	}
}

func (escore *Escore) AllGrade() ([]*Escore, error) {
	var escores []*Escore
	o := orm.NewOrm()
	_, err := o.QueryTable(Clazz{}).OrderBy("Id").All(&escores)
	if err != nil {
		return nil, err
	}
	return escores, nil
}

func (escore *Escore) Add() error {
	o := orm.NewOrm()
	if _, err := o.Insert(escore); err != nil {
		return err
	}
	return nil
}

func (escore *Escore) Update() error {
	o := orm.NewOrm()
	if _, err := o.Update(escore, "ExamId", "ClazzId", "StudentId", "GradeId", "CourseId", "Score"); err != nil {
		return err
	}
	return nil
}

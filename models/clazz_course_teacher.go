package models

import "github.com/astaxie/beego/orm"

type CCT struct {
	Id      int64
	Clazz   *Clazz   `orm:"rel(fk)"`
	Grade   *Grade   `orm:"rel(fk)"`
	Course  *Course  `orm:"rel(fk)"`
	Teacher *Teacher `orm:"rel(fk)"`
}

func (cct *CCT) OneGrade() {
	o := orm.NewOrm()
	err := o.Read(cct)
	if err != nil {
		return
	}
}

func (cct *CCT) AllGrade() ([]*CCT, error) {
	var ccts []*CCT
	o := orm.NewOrm()
	_, err := o.QueryTable(CCT{}).OrderBy("Id").All(&ccts)
	if err != nil {
		return nil, err
	}
	return ccts, nil
}

func (cct *CCT) Add() error {
	o := orm.NewOrm()
	if _, err := o.Insert(cct); err != nil {
		return err
	}
	return nil
}

func (cct *CCT) Update() error {
	o := orm.NewOrm()
	if _, err := o.Update(cct, "Id", "ClazzId", "GradeId", "CourseId", "TeacherId"); err != nil {
		return err
	}
	return nil
}

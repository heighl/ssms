package models

import "github.com/astaxie/beego/orm"

func (c *ClazzCourseTeacher) GetAll() ([]ClazzCourseTeacher, error) {
	var m []ClazzCourseTeacher
	o := orm.NewOrm()
	o.Using("default")
	qs:=o.QueryTable("clazz_course_teacher")
	_,err:=qs.All(&m)
	if err != nil {
		return nil, err
	}
	return m, nil
}


func GetClazz() ([]*Clazz, error) {
	var m []*Clazz
	_,err:= orm.NewOrm().QueryTable("clazz").Filter("Grade",1).RelatedSel().All(&m)
	if err!=nil{
		panic(err)
	}
	return m,nil
}
package models

import (
	"github.com/apex/log"
	"github.com/astaxie/beego/orm"
)

type Student struct {
	Id     int
	Number string
	Name   string
	Sex    string
	Phone  string
	Qq     string
	Clazz  *Clazz    `orm:"rel(fk)"`
	Grade  *Grade    `orm:"rel(fk)"`
	Escore []*Escore `orm:"reverse(many)"`
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

func (student *Student) AllStudent() ([]*Student, error) {
	var students []*Student
	o := orm.NewOrm()
	_, err := o.QueryTable(Student{}).OrderBy("Id").RelatedSel().All(&students)
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
	if _, err := o.Update(student, "Id", "Number", "Name", "Sex", "Phone", "Qq", "Clazz", "Grade"); err != nil {
		return err
	}
	return nil
}

//一共返回两个变量，一个是显示当前的。另外一个是没有分页的，可以很好的返回总页数
func (this *Student) ListLimit(limit, limit2, page int, key string) ([]*Student, []*Student) {
	o := orm.NewOrm()
	var clients []*Student
	var num []*Student
	if key == "*" {
		o.QueryTable(Student{}).Limit(limit, (page-1)*limit2).OrderBy("Id").RelatedSel().All(&clients)
		//log.Info(err.Error())
		o.QueryTable(Student{}).All(&num)
	} else {
		con := orm.NewCondition()
		con1 := con.Or("Name__icontains", key).Or("Phone__icontains", key)
		o.QueryTable(Student{}).SetCond(con1).Limit(limit, (page-1)*limit2).OrderBy("-Id").RelatedSel().All(&clients)
		o.QueryTable(Student{}).SetCond(con1).All(&num)

	}
	return clients, num
}
func (this *Student) AdressList(gradeId int64, clazzId int64) []*Student {
	var students []*Student
	o := orm.NewOrm()
	o.QueryTable(Student{}).Filter("Grade", gradeId).
		Filter("Clazz", clazzId).OrderBy("Id").RelatedSel().All(&students)
	return students
}

func (this *Student) NameGet(name string) []*Student {
	o := orm.NewOrm()
	var user []*Student
	_, err := o.QueryTable(Student{}).Filter("Name", name).All(&user)
	if err != nil {
		log.Info(err.Error())
	}
	return user
}

func (this *Student) GetId() *Student {
	o := orm.NewOrm()
	err := o.QueryTable(Student{}).Filter("Number", this.Number).RelatedSel().One(this)
	if err != nil {
		log.Error(err.Error())
	}
	return this
}

func (this *Student) DeletSt() {
	o := orm.NewOrm()
	_, err := o.Delete(this)
	if err != nil {
		log.Error(err.Error())
	}
}

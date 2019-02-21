package models

import (
	"github.com/apex/log"
	"github.com/astaxie/beego/orm"
)

type Escore struct {
	Id      int64
	Exam    *Exam    `orm:"rel(fk)"`
	Clazz   *Clazz   `orm:"rel(fk)"`
	Student *Student `orm:"rel(fk)"`
	Grade   *Grade   `orm:"rel(fk)"`
	Course  *Course  `orm:"rel(fk)"`
	Score   int64
}

func (this *Escore) GetOne() *Escore {
	o := orm.NewOrm()
	err := o.Read(this)
	if err != nil {
		log.Info(err.Error())
		return nil
	}
	return this
}

func (escore *Escore) AllGrade() ([]*Escore, error) {
	var escores []*Escore
	o := orm.NewOrm()
	_, err := o.QueryTable(Clazz{}).OrderBy("Id").RelatedSel().All(&escores)
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

func (escore *Escore) UpdateSorce() error {
	o := orm.NewOrm()
	if _, err := o.Update(escore,"Score"); err != nil {
		return err
	}
	return nil
}

//一共返回两个变量，一个是显示当前的。另外一个是没有分页的，可以很好的返回总页数
func (this *Escore)ListLimit(limit,page int,key string)([]*Escore,[]*Escore)  {
	o := orm.NewOrm()
	var clients []*Escore
	var num []*Escore
	if key=="*"{
		o.QueryTable(Escore{}).Limit(limit,(page-1)*limit).OrderBy("Id").RelatedSel().All(&clients)
		//log.Info(err.Error())
		o.QueryTable(Escore{}).All(&num)
	}else {
		con := orm.NewCondition()
		con1 := con.Or("Name__icontains",key).Or("Phone__icontains",key)
		o.QueryTable(Escore{}).SetCond(con1).Limit(limit,(page-1)*limit).OrderBy("-Id").RelatedSel().All(&clients)
		o.QueryTable(Escore{}).SetCond(con1).All(&num)

	}
	return clients,num
}

func (this *Escore) IdStudentEscore(id,types int) []*Escore  {
	o := orm.NewOrm()
	var students []*Escore
	o.QueryTable(Escore{}).Filter("Exam",types).Filter("Student",id).All(&students)
	return students
}
func (this *Escore) IdTypeEscore(id int) int {
	o:=orm.NewOrm()
	var students []*Escore
	o.QueryTable(Escore{}).Filter("Exam",id).All(&students)
	return len(students)/5
}
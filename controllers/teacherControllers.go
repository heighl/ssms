package controllers

import (
	"github.com/apex/log"
	"ssms/models"
)

type TeacherControllers struct {
	BaseController
}

func (this *TeacherControllers)List()  {
	teacher:=&models.Teacher{}
	teachers,err:=teacher.AllGrade()
	if err!=nil{
		log.Error(err.Error())
	}
	this.Data["teachers"]=teachers
	this.Layout="teacher/layout.html"
	this.TplName="teacher/getTeachers.html"
}
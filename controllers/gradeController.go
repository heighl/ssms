package controllers

import (
	"github.com/astaxie/beego"
	"ssms/models"
)

type GradeControllers struct {
	beego.Controller
}

func (this *GradeControllers)One()  {
	clazz:=&models.Clazz{Id:1}
	clazzid:=clazz.Grade()
	this.Data["json"]=&clazzid
	this.ServeJSON()
}

func (this *GradeControllers) AllList()  {
	grade:=&models.Clazz{}
	grades,err:=grade.AllGrade()
	if err!=nil{
		this.Ctx.WriteString("获取班级列表错误")
		this.StopRun()
	}
	this.Data["json"]=grades
	this.ServeJSON()
}
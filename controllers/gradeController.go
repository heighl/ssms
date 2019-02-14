package controllers

import (
	"github.com/astaxie/beego"
	"ssms/models"
)

type GradeControllers struct {
	beego.Controller
}

func (this *GradeControllers)One()  {
	grade:=models.Grade{Id:1}
	grade.OneGrade()
	this.Data["json"]=&grade
	this.ServeJSON()
}

func (this *GradeControllers) AllList()  {
	grade:=&models.Grade{}
	grades,err:=grade.AllGrade()
	if err!=nil{
		this.Ctx.WriteString("获取班级列表错误")
		this.StopRun()
	}
	this.Data["json"]=grades
	this.ServeJSON()
}
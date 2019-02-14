package controllers

import (
	"github.com/astaxie/beego"
	"ssms/models"
)

type StudentControllers struct {
	beego.Controller
}

func (this *StudentControllers)One()  {
	grade:=models.Student{Id:1}
	grade.OneGrade()
	this.Data["json"]=&grade
	this.ServeJSON()
}

func (this *StudentControllers) AllList()  {
	grade:=&models.Student{}
	grades,err:=grade.AllGrade()
	if err!=nil{
		this.Ctx.WriteString("获取班级列表错误")
		this.StopRun()
	}
	this.Data["json"]=grades
	this.ServeJSON()
}
package controllers

import (
	"github.com/astaxie/beego"
	"ssms/models"
)

type ExamControllers struct {
	beego.Controller
}

func (this *ExamControllers)One()  {
	grade:=models.Exam{Id:1}
	grade.OneGrade()
	this.Data["json"]=&grade
	this.ServeJSON()
}

func (this *ExamControllers) AllList()  {
	grade:=&models.Exam{}
	grades,err:=grade.AllGrade()
	if err!=nil{
		this.Ctx.WriteString("获取班级列表错误")
		this.StopRun()
	}
	this.Data["json"]=grades
	this.ServeJSON()
}
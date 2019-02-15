package controllers

import (
	"github.com/astaxie/beego"
	"ssms/models"
)

type GradeControllers struct {
	beego.Controller
}

func (st *GradeControllers)Get()  {
	clazz:=&models.Clazz{Id:1}
	g:=clazz.GetOne()

	st.Data["json"]=&g
	st.ServeJSON()
}

func (this *GradeControllers) AllList()  {
	var grade models.Clazz
	grades,err:=grade.AllGrade()
	if err!=nil{
		this.Ctx.WriteString("获取班级列表错误")
		this.StopRun()
	}
	this.Data["json"]=&grades
	this.ServeJSON()
}
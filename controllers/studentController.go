package controllers

import (
	"github.com/astaxie/beego"
	"ssms/models"
)

type StudentControllers struct {
	beego.Controller
}

func (st *StudentControllers) One() {
	clazz:=&models.Student{Id:1}
	g:=clazz.GetOne()

	st.Data["json"]=&g
	st.ServeJSON()
}

func (this *StudentControllers) AllList() {
	var grade models.Student
	grades, err := grade.AllGrade()
	if err != nil {
		this.Ctx.WriteString("获取班级列表错误")
		this.StopRun()
	}
	this.Data["json"] = &grades
	this.ServeJSON()
}

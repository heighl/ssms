package controllers

import (
	"github.com/astaxie/beego"
	"github.com/prometheus/common/log"
	"ssms/models"
)

type ClazzCourseTeacherController struct {
	beego.Controller
}

func (st *ClazzCourseTeacherController) Get()  {
	g,err:=models.GetClazz()
	if err!=nil{
		log.Info("getone err")
	}
	st.Data["json"]=&g
	st.ServeJSON()
}

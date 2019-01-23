package controllers

import (
	"github.com/astaxie/beego"
	"github.com/prometheus/common/log"
	"ssms/models"
)

type StudentController struct {
	beego.Controller
}

func (st *StudentController) Get()  {
	g,err:=models.GetOne(2)
	if err!=nil{
		log.Info("getone err")
	}
	st.Data["json"]=&g
	st.ServeJSON()
}
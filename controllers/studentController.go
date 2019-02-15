package controllers

import (
	"ssms/models"
)

type StudentControllers struct {
	BaseController
}

func (st *StudentControllers) One() {
	clazz:=&models.Student{Id:1}
	g:=clazz.GetOne()

	st.Data["json"]=&g
	st.ServeJSON()
}

func (this *StudentControllers)List()  {
	limit,err := this.GetInt("limit")
	if err !=nil{
		limit=10
	}
	page,err := this.GetInt("page")
	if err != nil{
		page=1
	}
	key := this.GetString("key")
	if key == ""{
		key = "*"
	}
	//如果提交的方式是搜索来的，必须定向到第一页
	if this.IsPost(){
		limit=10
		page=1
	}
	client := &models.Student{}
	clients,snum := client.ListLimit(limit,page,key)
	this.Data["students"]=clients
	this.Data["pagetitle"]="用户列表"
	//为了区分全搜索还是局部搜索要再次判断key
	if key == "*"{
		this.Data["key"]=""
	}else{
		this.Data["key"]=key
	}
	this.Data["pagecount"]=len(snum)
	this.Data["pagelimit"]=limit
	this.Data["page"]=page
	this.Xsrf()
	this.Layout="public/layout.html"
	this.TplName="studentcontroller/getStudent.html"
}
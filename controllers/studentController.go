package controllers

import (
	"github.com/apex/log"
	"ssms/models"
)

type StudentControllers struct {
	BaseController
}

func (this *StudentControllers) One() {
	account := this.GetSession("number")
	id := account.(string)

	student := models.Student{Number: id}
	st := student.GetId()
	this.Data["student"]=st
	this.Layout="studentcontroller/layout.html"
	this.TplName="studentcontroller/oneStudent.html"
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
	clients,snum := client.ListLimit(limit,limit,page,key)
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
	this.Layout="studentcontroller/layout.html"
	this.TplName="studentcontroller/getStudent.html"
}

func (this *StudentControllers)AddressList()  {
	id:=this.GetSession("number")
	ids:=id.(string)

	student:=models.Student{Number:ids}
	st:=student.GetId()

	students:=student.AdressList(st.Grade.Id,st.Clazz.Id)
	this.Data["students"]=students
	this.Layout="studentcontroller/layout.html"
	this.TplName="studentcontroller/getStudent.html"
}

func (this *StudentControllers)Updata()  {
	if this.IsPost() {
		//number:=this.GetString("number")
		tele := this.GetString("tele")
		qq := this.GetString("qq")
		account := this.GetSession("number")
		id := account.(string)

		student := models.Student{Number: id}
		st := student.GetId()
		st.Phone = tele
		st.Qq = qq

		err := st.Update()
		if err != nil {
			log.Error(err.Error())
		}
		log.Infof("stuent", st)

		//this.Data["student"]
		this.Redirect(this.URLFor(".One"),302)
	}else {
		this.Layout = "studentcontroller/layout.html"
		this.TplName = "studentcontroller/updateInformation.html"
	}
}
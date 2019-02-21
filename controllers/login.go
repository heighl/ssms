package controllers

import (
	"github.com/astaxie/beego"
	"ssms/libs"
	"ssms/models"
)

type LoginControllers struct {
	BaseController
}

func (this *LoginControllers)Login()  {
	if this.IsPost(){
		user := this.GetString("username")
		pass := this.GetString("password")
		password := libs.Passwords(pass)
		users := &models.User{Account:user,Password:password}
		if err := users.Login();err !=nil{
			this.Ctx.WriteString("验证失败")
			this.StopRun()
		}
		this.SetSession("id",users.Id)
		this.SetSession("number",users.Account)
		this.SetSession("type",users.Type)
		if  users.Type==1{
			this.Redirect(beego.URLFor("MainController.GetAdmin"),302)
		}else if users.Type==2 {
			this.Redirect(beego.URLFor("MainController.GetStudent"),302)
		}else {
			this.Redirect(beego.URLFor("MainController.GetTeacher"),302)
		}
	}else {
		this.Xsrf()
		this.Data["pagetitle"]="登录系统"
		this.TplName="index/index.html"
	}
}
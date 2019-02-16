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
		this.SetSession("username",user)
		this.Redirect(beego.URLFor("MainController.Get"),302)
	}else {
		this.Xsrf()
		this.Data["pagetitle"]="登录系统"
		this.TplName="index/index.html"
	}
}
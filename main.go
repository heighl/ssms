package main

import (
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"ssms/models"
	_ "ssms/routers"
)

func main() {
	models.Init()
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
}


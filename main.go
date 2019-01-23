package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"ssms/models"
	_ "ssms/routers"
	_ "github.com/go-sql-driver/mysql"
)

func init()  {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:123456@/ssms?charset=utf8")
	models.Init()
}

func main() {
	beego.Run()
}


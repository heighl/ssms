package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func Init() {
	//配置mysql
	host := beego.AppConfig.String("db.host")
	user := beego.AppConfig.String("db.user")
	password := beego.AppConfig.String("db.password")
	port := beego.AppConfig.String("db.port")
	dbName := beego.AppConfig.String("db.name")
	dsn := user + ":" + password + "@tcp" + "(" + host + ":" + port + ")/" + dbName + "?charset=utf8"
	orm.RegisterDriver("default", orm.DRMySQL)
	orm.RegisterDataBase("default", beego.AppConfig.String("db_adapter"), dsn)
	orm.RegisterModel(new(Grade), new(Clazz), new(Course), new(CCT), new(Escore), new(Exam), new(GradeCourse),
		new(Student), new(Teacher), new(User))
	beego.Info("数据库初始化完成.")
}

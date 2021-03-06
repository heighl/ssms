package routers

import (
	"github.com/astaxie/beego"
	"ssms/controllers"
)

func init() {
	//登录页面
	beego.Router("/login",&controllers.LoginControllers{},"*:Login")
	//学生列表
  	beego.Router("/student",&controllers.MainController{},"*:GetStudent")
	beego.Router("/student/info/list",&controllers.StudentControllers{},"*:List")
	beego.Router("/student/info/adresslist",&controllers.StudentControllers{},"*:AddressList")
	beego.Router("/student/info/update",&controllers.StudentControllers{},"*:Updata")
	beego.Router("/student/oneinfo",&controllers.StudentControllers{},"*:One")
  	//考试列表
	beego.Router("/exam/info/list",&controllers.ExamControllers{},"*:List")
	beego.Router("/exam/info",&controllers.GradeControllers{},"*:AllList")
  	//教师列表
	beego.Router("/teacher",&controllers.MainController{},"*:GetTeacher")
	beego.Router("/teacher/exam",&controllers.ExamControllers{},"*:TsList")
  	//
	//beego.Router("/clazz/info",&)
	//beego.Router("/grade/info",&)
	//beego.Router("/course/info",&)
	//beego.Router("/exam/info",&)
  	////学生成绩
	beego.Router("/escore",&controllers.EscoreControllers{},"*:List")
	beego.Router("/updataEscore",&controllers.EscoreControllers{},"*:EditEscore")
  	////教师登录成绩
	//beego.Router("/teacher/grade",&)
  	////考试统计
  	//beego.Router("/teacher/exam",&)
}

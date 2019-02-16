package routers

import (
	"github.com/astaxie/beego"
	"ssms/controllers"
)

func init() {
	//登录页面
	beego.Router("/login",&controllers.LoginControllers{},"*:Login")
	//学生列表
  	beego.Router("/student",&controllers.MainController{},"*:Get")
	beego.Router("/student/info/list",&controllers.StudentControllers{},"*:List")
	beego.Router("/student/info/adresslist",&controllers.StudentControllers{},"*:AddressList")
  	//考试列表
	beego.Router("/exam/info/list",&controllers.ExamControllers{},"*:List")
	beego.Router("/exam/info",&controllers.GradeControllers{},"*:AllList")
  	//教师列表
	//beego.Router("/teacher/info",&controllers.ClazzCourseTeacherController{})
  	////
	//beego.Router("/clazz/info",&)
	//beego.Router("/grade/info",&)
	//beego.Router("/course/info",&)
	//beego.Router("/exam/info",&)
  	////学生成绩
	beego.Router("/escore",&controllers.EscoreControllers{},"*:List")
  	////教师登录成绩
	//beego.Router("/teacher/grade",&)
  	////考试统计
  	//beego.Router("/teacher/exam",&)
}

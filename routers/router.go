package routers

import (
	"github.com/astaxie/beego"
	"ssms/controllers"
)

func init() {
	//学生列表
  	beego.Router("/student/info",&controllers.StudentControllers{},"*:One")
	beego.Router("/student/info/list",&controllers.StudentControllers{},"*:AllList")
  	//考试列表
	beego.Router("/exam/info",&controllers.GradeControllers{},"get:One")
	beego.Router("/exam/info/list",&controllers.GradeControllers{},"*:AllList")
  	//教师列表
	//beego.Router("/teacher/info",&controllers.ClazzCourseTeacherController{})
  	////
	//beego.Router("/clazz/info",&)
	//beego.Router("/grade/info",&)
	//beego.Router("/course/info",&)
	//beego.Router("/exam/info",&)
  	////学生成绩
	//beego.Router("/student/grade",&)
  	////教师登录成绩
	//beego.Router("/teacher/grade",&)
  	////考试统计
  	//beego.Router("/teacher/exam",&)
}

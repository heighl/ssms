package routers

import (
	"github.com/astaxie/beego"
	"ssms/controllers"
)

func init() {
	//学生列表
  	beego.Router("/student/info",&controllers.StudentController{})
  	//教师列表
	beego.Router("/teacher/info",&controllers.ClazzCourseTeacherController{})
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

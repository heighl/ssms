package controllers

import (
	"github.com/apex/log"
	"ssms/models"
	"strconv"
	"time"
)

type AdminControllers struct {
	BaseController
}

type ExamForm struct {
	Id     int64  `form:"-"`
	Name   string `form:"name"`
	Time   string `form:"time"`
	Remark string `form:"remark"`
	Type   string `form:"type"`
	Grade  string `form:"grade"`
	Clazz  string `form:"clazz"`
	Course string `form:"course"`
}

type StudentForm struct {
	Number string `form:"number"`
	Name   string `form:"name"`
	Sex    string `form:"sex"`
	Phone  string `form:"phone"`
	Qq     string `form:"qq"`
	Clazz  string `form:"clazz"`
	Grade  string `form:"grade"`
}

// ExamList 考试安排
func (this *AdminControllers) ExamList() {
	grade := &models.Exam{}
	grades, err := grade.AllGrade()
	if err != nil {
		this.Ctx.WriteString("获取班级列表错误")
		this.StopRun()
	}
	this.Data["exams"] = grades
	this.Layout = "admin/layout.html"
	this.TplName = "admin/exam.html"
}

// AddExam 添加考试
func (this *AdminControllers) AddExam() {
	if this.IsPost() {
		examform := ExamForm{}
		clazz := &models.Clazz{}
		if err := this.ParseForm(&examform); err != nil {
			log.Error(err.Error())
		}

		//解析年级
		formgrade, err := strconv.ParseInt(examform.Grade, 10, 64)
		if err != nil {
			log.Error(err.Error())
		}
		grade := &models.Grade{Id: formgrade}

		formcourse, err := strconv.ParseInt(examform.Course, 10, 64)
		if err != nil {
			log.Error(err.Error())
		}

		//解析班级，如果是0找到整个年级，否则通过班级名字和年级Id唯一确定一个班级。
		if examform.Clazz == "0" {
			formclazz, err := strconv.ParseInt(examform.Clazz, 10, 64)
			if err != nil {
				log.Error(err.Error())
			}
			clazz.Id = formclazz
			clazz = clazz.GetOne()
		} else {
			//在数据库找班级
			clazz.Name = examform.Clazz
			clazz.Grade = grade
			err = clazz.SerchClazz()
			if err != nil {
				log.Error(err.Error())
			}
		}

		formtype, err := strconv.ParseInt(examform.Type, 10, 64)
		if err != nil {
			log.Error(err.Error())
		}

		formtime, err := time.Parse("2006-01-02", examform.Time)
		if err != nil {
			log.Error(err.Error())
		}

		course := &models.Course{Id: formcourse}
		exam := &models.Exam{
			Id:     examform.Id,
			Name:   examform.Name,
			Time:   formtime,
			Remark: examform.Remark,
			Type:   formtype,
			Grade:  grade.OneGrade(),
			Clazz:  clazz,
			Course: course.OneCourse(),
		}

		err = exam.Add()
		if err != nil {
			log.Error(err.Error())
		}
		this.Redirect(this.URLFor(".ExamList"), 302)
	} else {
		this.Layout = "admin/layout.html"
		this.TplName = "admin/addExam.html"
	}
}

// ExamList 学生信息
func (this *AdminControllers) StudentInfoList() {
	limit, err := this.GetInt("limit")
	if err != nil {
		limit = 10
	}
	page, err := this.GetInt("page")
	if err != nil {
		page = 1
	}
	key := this.GetString("key")
	if key == "" {
		key = "*"
	}

	//如果提交的方式是搜索来的，必须定向到第一页
	if this.IsPost() {
		limit = 10
		page = 1
	}
	client := &models.Student{}
	clients, snum := client.ListLimit(limit, limit, page, key)
	this.Data["students"] = clients
	this.Data["pagetitle"] = "用户列表"
	//为了区分全搜索还是局部搜索要再次判断key
	if key == "*" {
		this.Data["key"] = ""
	} else {
		this.Data["key"] = key
	}
	this.Data["pagecount"] = len(snum)
	this.Data["pagelimit"] = limit
	this.Data["page"] = page
	this.Xsrf()
	this.Layout = "admin/layout.html"
	this.TplName = "admin/studentInfo.html"
}

// AddStudentInfo 添加学生信息
func (this *AdminControllers) AddStudentInfo() {
	if this.IsPost() {
		clazz := &models.Clazz{}
		studnetform := StudentForm{}
		if err := this.ParseForm(&studnetform); err != nil {
			log.Error(err.Error())
		}

		//解析年级
		formgrade, err := strconv.ParseInt(studnetform.Grade, 10, 64)
		if err != nil {
			log.Error(err.Error())
		}
		grade := &models.Grade{Id: formgrade}

		//在数据库找班级
		clazz.Name = studnetform.Clazz
		clazz.Grade = grade
		err = clazz.SerchClazz()
		if err != nil {
			log.Error(err.Error())
		}

		student := &models.Student{
			Number: studnetform.Number,
			Name:   studnetform.Name,
			Sex:    studnetform.Sex,
			Phone:  studnetform.Phone,
			Qq:     studnetform.Qq,
			Clazz:  clazz,
			Grade:  grade,
		}

		err = student.Add()
		if err != nil {
			log.Error(err.Error())
		}
		this.Redirect(this.URLFor(".StudentInfoList"), 302)
	}

	this.Layout = "admin/layout.html"
	this.TplName = "admin/addStudentInfo.html"
}

// ExamList 修改学生信息
func (this *AdminControllers) UpStudentInfoList() {
	limit, err := this.GetInt("limit")
	if err != nil {
		limit = 10
	}
	page, err := this.GetInt("page")
	if err != nil {
		page = 1
	}
	key := this.GetString("key")
	if key == "" {
		key = "*"
	}

	//如果提交的方式是搜索来的，必须定向到第一页
	if this.IsPost() {
		limit = 10
		page = 1
	}
	client := &models.Student{}
	clients, snum := client.ListLimit(limit, limit, page, key)
	this.Data["students"] = clients
	this.Data["pagetitle"] = "用户列表"
	//为了区分全搜索还是局部搜索要再次判断key
	if key == "*" {
		this.Data["key"] = ""
	} else {
		this.Data["key"] = key
	}
	this.Data["pagecount"] = len(snum)
	this.Data["pagelimit"] = limit
	this.Data["page"] = page
	this.Xsrf()
	this.Layout = "admin/layout.html"
	this.TplName = "admin/upStudentList.html"
}

func (this *AdminControllers) UpdataStudentInfo() {
	if this.IsPost() {
		clazz := &models.Clazz{}
		id, err := this.GetInt("id")
		if err != nil {
			log.Info(err.Error())
		}
		grade := this.GetString("grade")
		formClazz := this.GetString("clazz")
		//解析年级
		formgrade, err := strconv.ParseInt(grade, 10, 64)
		if err != nil {
			log.Error(err.Error())
		}
		newGrade := &models.Grade{Id: formgrade}

		//在数据库找班级
		clazz.Name = formClazz
		clazz.Grade = newGrade
		err = clazz.SerchClazz()
		if err != nil {
			log.Error(err.Error())
		}

		//修改信息
		student := &models.Student{Id: id}
		student.GetOne()
		number := this.GetString("number")
		if number != "" {
			student.Number = number
		}

		name := this.GetString("name")
		if name != "" {
			student.Name = name
		}

		sex := this.GetString("sex")
		if sex != "" {
			student.Sex = sex
		}

		tele := this.GetString("tele")
		if tele != "" {
			student.Phone = tele
		}

		qq := this.GetString("qq")
		if qq != "" {
			student.Qq = qq
		}

		if formClazz != "" {
			student.Clazz = clazz
		}
		if grade != "" {
			student.Grade = newGrade
		}

		err = student.Update()
		if err != nil {
			log.Info("修改学生信息错误")
		}
		this.Redirect(this.URLFor(".UpStudentInfoList"), 302)
	}
	id, err := this.GetInt("id")
	if err != nil {
		log.Info(err.Error())
	}
	this.Data["id"] = id
	this.Layout = "admin/layout.html"
	this.TplName = "admin/updataStudent.html"
}

// DeleteStudent 删除学生页面
func (this *AdminControllers) DeleteStudentPage() {
	limit, err := this.GetInt("limit")
	if err != nil {
		limit = 10
	}
	page, err := this.GetInt("page")
	if err != nil {
		page = 1
	}
	key := this.GetString("key")
	if key == "" {
		key = "*"
	}

	//如果提交的方式是搜索来的，必须定向到第一页
	if this.IsPost() {
		limit = 10
		page = 1
	}
	client := &models.Student{}
	clients, snum := client.ListLimit(limit, limit, page, key)
	this.Data["students"] = clients
	this.Data["pagetitle"] = "用户列表"
	//为了区分全搜索还是局部搜索要再次判断key
	if key == "*" {
		this.Data["key"] = ""
	} else {
		this.Data["key"] = key
	}
	this.Data["pagecount"] = len(snum)
	this.Data["pagelimit"] = limit
	this.Data["page"] = page
	this.Xsrf()
	this.Layout = "admin/layout.html"
	this.TplName = "admin/deleteStudentList.html"
}

func (this *AdminControllers) DeleteSt() {
	id, err := this.GetInt("id")
	if err != nil {
		log.Info(err.Error())
	}
	student := &models.Student{Id: id}
	student.DeletSt()
	this.Redirect(this.URLFor(".DeleteStudentPage"), 302)
}

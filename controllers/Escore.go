package controllers

import (
	"github.com/apex/log"
	"github.com/astaxie/beego"
	"ssms/models"
	"strconv"
	"strings"
)

type EscoreControllers struct {
	BaseController
}

func (st *EscoreControllers) Get() {
	clazz := &models.Escore{Id: 1}
	g := clazz.GetOne()

	st.Data["json"] = &g
}

func (this *EscoreControllers) List() {
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
	types,err := this.GetInt("id")
	if err!=nil{
		log.Info(err.Error())
	}

	//如果提交的方式是搜索来的，必须定向到第一页
	if this.IsPost() {
		limit = 10
		page = 1
	}
	escore:=new(models.Escore)
	count:=escore.IdTypeEscore(types)
	this.Data["pagecount"] = count

	client := &models.Student{}
	var clients []*models.Student
	if count/limit>=page{
		clients, _ = client.ListLimit(limit,limit, page, key)
	}else {
		clients, _ = client.ListLimit(count%limit,limit, page, key)
	}
	//clients, _ := client.ListLimit(limit, page, key)
	//有没有更好的方式，在不重组数据的情况下，直接通过orm获得所需要的值？这点太不方便了。
	list := make([]map[string]interface{}, len(clients))
	for k, v := range clients {
		if v.Id<=count{
			row := make(map[string]interface{})
			row["Number"] = v.Number
			row["Name"] = v.Name

			escore := new(models.Escore)
			escores := escore.IdStudentEscore(v.Id,types)
			var escoresid []string
			for i:=0;i<len(escores);i++  {
				escoresid=append(escoresid, strconv.FormatInt(escores[i].Id,10))
			}
			row["escoresid"]=escoresid
			row["Chinese"]=escores[0].Score
			row["Math"]=escores[1].Score
			row["English"]=escores[2].Score
			row["Py"]=escores[3].Score
			row["Bi"]=escores[4].Score
			var a int64
			for i:=0;i<len(escores);i++ {
				a+=escores[i].Score
			}
			row["Total"] = a
			list[k] = row
		}
	}
	this.Data["list"] = list
	this.Data["pagetitle"] = "成绩列表"
	//为了区分全搜索还是局部搜索要再次判断key
	if key == "*" {
		this.Data["key"] = ""
	} else {
		this.Data["key"] = key
	}

	this.Data["id"]=types
	this.Data["pagelimit"] = limit
	this.Data["page"] = page
	this.Xsrf()
	usertype:=this.GetSession("type")
	usert:=usertype.(int64)
	this.Data["type"]=usert
	if usert==3{
		this.SetSession("ExamId",types)
		this.Layout = "teacher/layout.html"
		this.TplName = "teacher/escore.html"
	}else if usert==2{
		this.Layout = "studentcontroller/layout.html"
		this.TplName = "studentcontroller/escore.html"
	}
}

func (this *EscoreControllers)EditEscore()  {
	if this.IsPost(){
		chinese,err:=this.GetInt64("Chinese")
		if err!=nil{
			log.Error(err.Error())
		}
		math,err:=this.GetInt64("Math")
		if err!=nil{
			log.Error(err.Error())
		}
		english,err:=this.GetInt64("English")
		if err!=nil{
			log.Error(err.Error())
		}
		py,err:=this.GetInt64("Py")
		if err!=nil{
			log.Error(err.Error())
		}
		bi,err:=this.GetInt64("Bi")
		if err!=nil{
			log.Error(err.Error())
		}
		//得到考试类型
		id:=this.GetSession("ExamId")
		studentid:=this.GetString("escoresid")
		studentid=strings.Replace(studentid,"[","",1)
		studentid=strings.Replace(studentid,"]","",1)
		st:=strings.Split(studentid," ")

		UpdataScore(chinese,math,english,py,bi,st)
		this.Layout = "teacher/layout.html"
		this.Redirect(beego.URLFor("EscoreControllers.List","id",id),302)
	}else {
		number:=this.GetString("number")
		this.Data["number"]=number
		this.Layout = "teacher/layout.html"
		this.TplName = "teacher/updataEscore.html"
	}
}

func UpdataScore(chinese,math,english,py,bi int64,studentid []string)  {
	for i:=0;i<len(studentid);i++ {
		esid,err:=strconv.ParseInt(studentid[i],10,64)
		if err!=nil{
			log.Error(err.Error())
		}
		escore:=&models.Escore{Id:esid}
		if esid%5==0{
			escore.Score=chinese
			escore.UpdateSorce()
		}else if esid%5==1{
			escore.Score=math
			escore.UpdateSorce()
		} else if esid%5==2{
			escore.Score=english
			escore.UpdateSorce()
		} else if esid%5==3{
			escore.Score=py
			escore.UpdateSorce()
		} else {
			escore.Score=bi
			escore.UpdateSorce()
		}
	}

}
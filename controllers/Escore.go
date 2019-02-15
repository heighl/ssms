package controllers

import (
	"ssms/models"
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
	//如果提交的方式是搜索来的，必须定向到第一页
	if this.IsPost() {
		limit = 10
		page = 1
	}
	escore:=new(models.Escore)
	count:=escore.IdTypeEscore(1)
	this.Data["pagecount"] = count

	client := &models.Student{}
	var clients []*models.Student
	if count/limit>=page{
		clients, _ = client.ListLimit(limit, page, key)
	}else {
		clients, _ = client.ListLimit(count%limit, page, key)
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
			escores := escore.IdStudentEscore(v.Id)
			row["Chinese"]=escores[0].Score
			row["Math"]=escores[1].Score
			row["English"]=escores[2].Score
			row["Py"]=escores[3].Score
			row["Bi"]=escores[4].Score
			var a int64
			for i:=0;i<len(escores)-1;i++ {
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

	this.Data["pagelimit"] = limit
	this.Data["page"] = page
	this.Xsrf()
	this.Layout = "studentcontroller/layout.html"
	this.TplName = "studentcontroller/escore.html"
}

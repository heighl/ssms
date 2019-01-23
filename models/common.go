package models

import "github.com/astaxie/beego/orm"

func GetOne(i int64) (Grade, error) {
	m:=Grade{Id:i}
	o := orm.NewOrm()
	o.Using("default")
	err:=o.Read(&m)
	if err != nil {
		return m,err
	}
	return m,nil
}


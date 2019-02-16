package models

import (
	"github.com/apex/log"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id       int64
	Account  string
	Password string
	Name     string
	Type     int64
}

func (this *User) GetOne() *User {
	o := orm.NewOrm()
	err := o.Read(this)
	if err != nil {
		log.Info(err.Error())
		return nil
	}
	return this
}

func (this *User) AllGrade() ([]*User, error) {
	var users []*User
	o := orm.NewOrm()
	_, err := o.QueryTable(User{}).OrderBy("Id").All(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (this *User) Add() error {
	o := orm.NewOrm()
	if _, err := o.Insert(this); err != nil {
		return err
	}
	return nil
}

func (this *User) Update() error {
	o := orm.NewOrm()
	if _, err := o.Update(this, "Name"); err != nil {
		return err
	}
	return nil
}

func (this *User)Login()error  {
	o := orm.NewOrm()
	if err :=o.Read(this,"Account");err !=nil{
		log.Info(err.Error())
		return err
	}
	return nil
}

func (this *User) IdGet(id interface{}) []*User {
	o := orm.NewOrm()
	var user []*User
	_,err:=o.QueryTable(User{}).Filter("Account",id).All(&user)
	if err!=nil{
		log.Info(err.Error())
	}
	return user
}
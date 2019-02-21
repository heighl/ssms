package controllers

type MainController struct {
	BaseController
}

func (this *MainController) GetStudent() {

	this.TplName="studentcontroller/layout.html"


}

func (this *MainController) GetTeacher() {

	this.TplName="teacher/layout.html"

}

func (this *MainController) GetAdmin() {

	this.TplName="admin/layout.html"

}
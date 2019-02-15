package controllers

type MainController struct {
	BaseController
}

func (this *MainController) Get() {

	this.TplName="studentcontroller/layout.html"


}

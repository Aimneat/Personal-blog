package controllers

import (
	"beeblog/models"
	"github.com/astaxie/beego/logs"

	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Get() {
	this.Data["IsHome"] = true
	this.TplName = "home.html"
	this.Data["IsLogin"] = checkAccount(this.Ctx)

	topics, err := models.GetAllTopics(
		this.Input().Get("cate"), this.Input().Get("lable"), true)
	if err != nil {
		logs.Error(err)
	}
	this.Data["Topics"] = topics

	categories, err := models.GetAllCategories()
	if err != nil {
		logs.Error(err)
	}
	this.Data["Categories"] = categories
}

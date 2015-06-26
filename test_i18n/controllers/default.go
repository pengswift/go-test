package controllers

import (
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
)

type baseController struct {
	beego.Controller
	i18n.Locale
}

func (this *baseController) Prepare() {
	lang := this.GetString("lang")
	if lang == "zh-CN" {
		this.Lang = lang
	} else {
		this.Lang = "en-US"
	}
	this.Data["Lang"] = this.Lang
}

type MainController struct {
	baseController
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["Hi"] = c.Tr("hi")
	c.Data["Bye"] = c.Tr("bye")
	c.Data["About"] = "about"
	c.Data["Sex"] = ".sex."
	c.TplNames = "index.tpl"
}

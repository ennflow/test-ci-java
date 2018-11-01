package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "www.enncloud.cn"
	c.Data["Email"] = "hr@ennew.cn"
	c.TplName = "index.tpl"
}

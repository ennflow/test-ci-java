package main

import (
	"beeblog/controllers"
	_"beeblog/models"
	"github.com/astaxie/beego"
	_"github.com/astaxie/beego/orm"
)


func main() {

	beego.Router("/", &controllers.MainController{})
	beego.Run()
}

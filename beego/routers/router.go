package routers

import (
	"github.com/weifabing/learngo/beego/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}

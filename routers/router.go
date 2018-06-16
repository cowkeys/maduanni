package routers

import (
	"github.com/cowkeys/maduanni/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}

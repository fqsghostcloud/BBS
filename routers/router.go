package routers

import (
	"bbs/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/signup", &controllers.SignupController{})
	beego.Router("/login", &controllers.LoginController{})
}

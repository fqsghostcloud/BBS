package routers

import (
	"bbs/controllers"
	"bbs/controllers/user"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &user.LoginController{})
	beego.Router("/signup", &user.SignupController{})
}

package routers

import (
	"bbs/controllers/admin"
	"bbs/controllers/user"

	"github.com/astaxie/beego"
)

func init() {
	beego.Include(&user.LoginController{})
	beego.Include(&user.SignupController{})
	beego.Include(&user.ManageController{})
	beego.Include(&admin.ManageController{})
}

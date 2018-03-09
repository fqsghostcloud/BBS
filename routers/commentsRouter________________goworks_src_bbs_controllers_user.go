package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["bbs/controllers/user:LoginController"] = append(beego.GlobalControllerRouter["bbs/controllers/user:LoginController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/user/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bbs/controllers/user:ManageController"] = append(beego.GlobalControllerRouter["bbs/controllers/user:ManageController"],
		beego.ControllerComments{
			Method: "SearchUser",
			Router: `/user/search`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bbs/controllers/user:SignupController"] = append(beego.GlobalControllerRouter["bbs/controllers/user:SignupController"],
		beego.ControllerComments{
			Method: "ActiveAccount",
			Router: `/user/auth/?:token`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bbs/controllers/user:SignupController"] = append(beego.GlobalControllerRouter["bbs/controllers/user:SignupController"],
		beego.ControllerComments{
			Method: "Signup",
			Router: `/user/signup`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}

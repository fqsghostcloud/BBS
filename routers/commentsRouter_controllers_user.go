package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["bbs/controllers/user:LoginController"] = append(beego.GlobalControllerRouter["bbs/controllers/user:LoginController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bbs/controllers/user:SignupController"] = append(beego.GlobalControllerRouter["bbs/controllers/user:SignupController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/signup`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}

package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["bbs/controllers/admin:ManageController"] = append(beego.GlobalControllerRouter["bbs/controllers/admin:ManageController"],
		beego.ControllerComments{
			Method: "Active",
			Router: `/admin/activeuser`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bbs/controllers/admin:ManageController"] = append(beego.GlobalControllerRouter["bbs/controllers/admin:ManageController"],
		beego.ControllerComments{
			Method: "Deactive",
			Router: `/admin/deactiveuser`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bbs/controllers/admin:ManageController"] = append(beego.GlobalControllerRouter["bbs/controllers/admin:ManageController"],
		beego.ControllerComments{
			Method: "DeleteUser",
			Router: `/admin/delete`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}

package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["bbs/controllers/admin:ManageController"] = append(beego.GlobalControllerRouter["bbs/controllers/admin:ManageController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/admin/delete`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bbs/controllers/admin:ManageController"] = append(beego.GlobalControllerRouter["bbs/controllers/admin:ManageController"],
		beego.ControllerComments{
			Method: "FuzzySearch",
			Router: `/admin/fsearch`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bbs/controllers/admin:ManageController"] = append(beego.GlobalControllerRouter["bbs/controllers/admin:ManageController"],
		beego.ControllerComments{
			Method: "Search",
			Router: `/admin/search`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}

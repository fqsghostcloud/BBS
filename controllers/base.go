package controllers

import (
	"net/http"

	"github.com/astaxie/beego"
)

type baseController struct {
	beego.Controller
}

// if everything is ok ,record to do what and result
func (b *baseController) serverOk(msg map[string]interface{}) {
	if len(msg) > 0 {
		b.Data["json"] = msg
	}
	c.Ctx.ResponseWriter.WriteHeader(http.StatusOK)
	b.ServeJSON()
}

// has error, record error msg , http code
func (b *baseController) serverError(msg map[string]interface{}, code int) {
	data["code"] = code

	if len(msg) > 0 {
		data["error"] = msg
	}
	c.Ctx.ResponseWriter.WriteHeader(code)
	b.ServeJSON()
}

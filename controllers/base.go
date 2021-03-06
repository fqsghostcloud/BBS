package controllers

import (
	"bbs/models"
	"net/http"

	"github.com/astaxie/beego"
)

// BaseController .
type BaseController struct {
	beego.Controller
}

// Manager all model manager
var Manager models.Manager

func init() {
	Manager = models.NewManager()
}

// ServerOk if everything is ok ,record to do what and result
func (b *BaseController) ServerOk(msg map[string]interface{}) {
	b.Data["json"] = msg // set msg to clinet json
	b.Ctx.ResponseWriter.WriteHeader(http.StatusOK)
	b.ServeJSON()
}

// ServerError has error, record error msg , http code
func (b *BaseController) ServerError(msg map[string]interface{}, code int) {
	msg["code"] = code
	b.Data["json"] = msg // se msg to client json
	b.Ctx.ResponseWriter.WriteHeader(code)
	b.ServeJSON()
}

package controllers

import (
	"net/http"
)


type LoginController struct {
	baseController
}

func (c *LoginController) Get() {

}

func (c *LoginController) Post() {
	username := c.GetString("username")
	password := c.GetString("password")

	data := map[string]string

	if  status, err := database.IsExsit("username") {
		data["info"] = "此用户名不存在"
		c.serverError(data, http.StatusBadRequest)
		return
	}


}

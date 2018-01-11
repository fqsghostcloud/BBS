package controllers

import (
	"fmt"
	"net/http"
)

type SignupController struct {
	baseController
}

func (c *SignupController) Get() {

}

func (c *SignupController) Post() {
	username := c.GetString("username")
	password := c.GetString("password")

	// set http header for json

	data := map[string]interface{}{}

	if database.IsExsit(username) {
		data["error"] = "该用户名已经存在"
		c.serverError(data, http.StatusBadRequest)
		return
	}

	if status, err := database.Signup(username, password); err != nil {
		fmt.Println(err)
	}

	if !status {
		data["error"] = "注册失败，请联系管理员"
		c.serverError(data, http.StatusBadRequest)
		return
	}

	data["info"] = "注册成功"
	c.serverOk(data)
	return

}

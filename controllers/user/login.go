package user

import (
	"bbs/controllers"
	"bbs/models/types"
	"bbs/models/user"
	"net/http"

	"github.com/astaxie/beego/orm"

	"github.com/golang/glog"
)

type LoginController struct {
	controllers.BaseController
}

func (c *LoginController) Get() {

}

// @router /login [post]
func (c *LoginController) Post() {
	username := c.GetString("username")
	password := c.GetString("password")

	data := map[string]interface{}{}
	dbUser := user.User{}

	ok, err := dbUser.Auth(username, password)
	if err != nil {
		if err == orm.ErrNoRows {
			data["info"] = "用户不存在"
			c.ServerOk(data)
		} else if err.Error() == types.UserLogForbidden {
			data["info"] = types.UserLogForbidden
			c.ServerOk(data)
		} else {
			glog.Errorf("auth username[%s], password[%s], error[%s]\n", username, password, err.Error())
			data["error"] = "登录过程中发生错误, 登录失败!"
			c.ServerError(data, http.StatusBadRequest)
		}

		return
	}

	if ok {
		glog.Infof("login success username[%s]\n", username)
		data["info"] = "登录成功"
		// use session
		c.SetSession("isLogin", true)
		c.SetSession("username", username)

	} else {
		glog.Infof("login faild username[%s], password[%s]\n", username, password)
		data["info"] = "登录失败,密码错误"
	}
	c.ServerOk(data)
	return
}

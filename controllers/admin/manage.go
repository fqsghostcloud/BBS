package admin

import (
	"bbs/controllers"
	"bbs/models/types"
	"bbs/models/user"
	"fmt"
	"net/http"

	"github.com/golang/glog"
)

// ManageController for admin or root manage bbs
type ManageController struct {
	controllers.BaseController
}

func (m *ManageController) Get() {

}

//DeleteUser user
// @router /admin/delete [post]
func (m *ManageController) DeleteUser() {
	username := m.GetString("username")
	data := map[string]interface{}{}

	err := controllers.Manager.DelUser(&user.User{Name: username})
	if err != nil {
		if err.Error() == types.UserNotExsit {
			data["info"] = err.Error()
			m.ServerOk(data)
			glog.Infof("del user info[%s]", err.Error())
		} else {
			data["error"] = err.Error()
			m.ServerError(data, http.StatusInternalServerError)
			glog.Errorf("del user error[%s]", err.Error())
		}
		return
	}

	data["info"] = "删除用户成功"
	m.ServerOk(data)
	glog.Infof("delete user[%s] success\n", username)
	return
}

// Active active user
// @router /admin/activeuser [post]
func (m *ManageController) Active() {
	username := m.GetString("username")
	data := map[string]interface{}{}

	err := controllers.Manager.ActiveUser(username)
	if err != nil {
		data["error"] = err.Error()
		glog.Errorf("active user error:[%s]", err.Error())
		m.ServerError(data, http.StatusBadRequest)
		return
	}

	data["info"] = fmt.Sprintf("激活用户[%s]成功\n", username)
	m.ServerOk(data)
	return
}

// Deactive inactive user
// @router /admin/deactiveuser [post]
func (m *ManageController) Deactive() {
	username := m.GetString("username")
	data := map[string]interface{}{}

	err := controllers.Manager.DeactiveUser(username)
	if err != nil {
		data["error"] = err.Error()
		glog.Errorf("inactive user[%s] error[%s]", username, err.Error())
		m.ServerError(data, http.StatusBadRequest)
		return
	}

	data["info"] = fmt.Sprintf("冻结用户[%s]成功\n", username)
	glog.Infoln(data["info"])
	m.ServerOk(data)
	return
}

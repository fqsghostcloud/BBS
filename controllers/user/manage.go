package user

import (
	"bbs/controllers"
	"bbs/models/types"
	"bbs/models/user"
	"net/http"

	"github.com/golang/glog"
)

type ManageController struct {
	controllers.BaseController
}


//Get user info 精准查找/模糊查找
func (m *ManageController) Get() {
	username := m.GetString("username")

	dbUser := user.User{}
	data := map[string]int{}{}
}

// Put modify user Info
func (m *ManageController) Put() {
	username := m.GetString("username")

	dbUser := user.User{}
	//用户信息不全????
	dbUser.Name = username
	data := map[string]interface{}{}

	err := dbUser.UpdateUser(&dbUser)
	if err != nil {
		if err.Error() == types.UsernameExErr {
			data["info"] = err.Error()
			m.ServerOk(data)
			glog.Infof("modify user info[%s]", err.Error())
		} else {
			data["error"] = err.Error()
			m.ServerError(data, http.StatusInternalServerError)
			glog.Errorf("modify user error[%s]", err.Error())
		}
		return
	}
}

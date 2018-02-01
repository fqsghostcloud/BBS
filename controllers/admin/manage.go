package admin

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

func (m *ManageController) Get() {

}

// Delete User
func (m *ManageController) Delete() {
	username := m.GetString("username")

	data := map[string]interface{}{}
	dbUser := user.User{}

	dbUser.Name = username

	err := dbUser.DelUser(&dbUser)
	if err != nil {
		if err.Error() == types.UsernameExErr {
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
}

package admin

import (
	"bbs/controllers"
	"bbs/models/types"
	"bbs/models/user"
	"net/http"

	"github.com/golang/glog"
)

// ManageController for admin or root manage bbs
type ManageController struct {
	controllers.BaseController
}

func (m *ManageController) Get() {

}

//Delete user
// @router /admin/delete [post]
func (m *ManageController) Delete() {
	username := m.GetString("username")

	data := map[string]interface{}{}
	dbUser := user.User{}

	dbUser.Name = username

	err := dbUser.DelUser(&dbUser)
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

// FuzzySearch search user by fuzzy name
// @router /admin/fsearch [post]
func (m *ManageController) FuzzySearch() {
	username := m.GetString("username")
	data := map[string]interface{}{}
	dbUser := user.User{}

	users, err := dbUser.FuzzySearch(username)
	if err != nil {
		glog.Errorf("fuzzy search error:[%s]", err.Error())
		data["error"] = types.Error
		m.ServerError(data, http.StatusBadRequest)
		return
	}

	if len(*users) == 0 {
		data["info"] = types.DataNotExsit
		m.ServerOk(data)
		return
	}

	data["info"] = users
	m.ServerOk(data)
	return

}

// Search search user info
// @router /admin/search [post]
func (m *ManageController) Search() {
	username := m.GetString("username")
	data := map[string]interface{}{}
	dbUser := user.User{}

	userInfo, err := dbUser.Search(username)
	if err != nil {
		if err.Error() == types.RowNotFound {
			data["info"] = types.DataNotExsit
			m.ServerOk(data)
		} else {
			glog.Errorf("search user error:[%s]", err.Error())
			data["error"] = types.Error
			m.ServerError(data, http.StatusBadRequest)
		}
		return
	}

	data["info"] = userInfo
	m.ServerOk(data)
	return
}

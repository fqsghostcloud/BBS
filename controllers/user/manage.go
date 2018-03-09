package user

import (
	"bbs/controllers"
	"bbs/models/types"
	"bbs/models/user"
	"net/http"

	"github.com/golang/glog"
)

// ManageController .
type ManageController struct {
	controllers.BaseController
}

//SearchUser user info 精准查找/模糊查找
// @router /user/search [post]
func (m *ManageController) SearchUser() {
	littleName := m.GetString("littleName")
	allName := m.GetString("allName")
	data := map[string]interface{}{}
	if len(littleName) == 0 && len(allName) == 0 {
		data["info"] = "搜索名称不能为空"
		m.ServerOk(data)
		return
	}

	if len(littleName) > 0 && len(allName) > 0 {
		data["info"] = "精准/模糊查找，不可以同时进行"
		m.ServerOk(data)
		return
	}

	if len(littleName) > 0 {
		users, error := controllers.Manager.FuzzySearch(littleName)
		if error != nil {
			data["error"] = error.Error()
			glog.Errorf("Fuzzy search error:[%s]\n", error.Error())
			m.ServerError(data, http.StatusBadRequest)
			return
		}

		data["resp"] = users
	}

	if len(allName) > 0 {
		user, err := controllers.Manager.Search(allName)
		if err != nil {
			data["error"] = err.Error()
			glog.Errorf("search user error:[%s]\n", err.Error())
			m.ServerError(data, http.StatusBadRequest)
			return
		}

		data["resp"] = user
	}

	m.ServerOk(data)
	return
}

// Put modify user Info
func (m *ManageController) Put() {
	username := m.GetString("username")

	//用户信息不全????
	data := map[string]interface{}{}

	err := controllers.Manager.UpdateUser(&user.User{Name: username})
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

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
	// username := m.GetString("username")

	// dbUser := user.User{}
	// data := map[string]interface{}{}
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

// var authEmail *user.Email

// // AuthEmailURL auth user email active url
// // @router /user/auth/emailurl [post]
// func (m *ManageController) AuthEmailURL() { // 此结构仅测试功能
// 	url := m.GetString("url")
// 	data := map[string]interface{}{}
// 	if authEmail != nil {
// 		status, emial := authEmail.CheckEmailURL(url)
// 		data["Info"] = fmt.Sprintf("Auth email[%s] status[%t]", emial, status)
// 		m.ServerOk(data)
// 		authEmail = nil
// 		return
// 	}
// 	err := fmt.Errorf("AuthEmail already nil")
// 	data["error"] = fmt.Sprintf("Auth emial url error[%s]", err.Error())
// 	m.ServerError(data, http.StatusBadRequest)
// 	return

// }

// // GenerateAuthEmailURL ..
// // @router /user/auth/email [post]
// func (m *ManageController) GenerateAuthEmailURL() { // 此接口仅测试功能
// 	email := m.GetString("email")
// 	authEmail = new(user.Email)
// 	authURL := authEmail.GenerateAuthURL(email)

// 	data := map[string]interface{}{}
// 	data["url"] = authURL
// 	m.ServerOk(data)
// 	return
// }

// //SendActivationURL to auth email and acitve user
// // @router /user/send/acemail [post]
// func (m *ManageController) SendActivationURL() { // 此接口仅测试功能
// 	email := user.Email{}
// 	data := map[string]interface{}{}

// 	email.SetTheme("帐号激活")
// 	email.SetEmailContent("邮箱内容")

// 	err := email.InitSendCfg(m.GetString("email"), m.GetString("username"))
// 	if err != nil {
// 		glog.Errorf("send emial init config error: [%s]", err.Error())
// 		data["error"] = err.Error()
// 		m.ServerError(data, http.StatusInternalServerError)
// 		return
// 	}

// 	err = email.SendEmail()
// 	if err != nil {
// 		data["error"] = err.Error()
// 		m.ServerError(data, http.StatusInternalServerError)
// 		return
// 	}

// 	data["info"] = fmt.Sprintf("send email to [%s] success!", m.GetString("email"))
// 	m.ServerOk(data)
// 	return
// }

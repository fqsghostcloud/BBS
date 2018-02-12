package user

import (
	"bbs/controllers"
	"bbs/models/types"
	"bbs/models/user"
	"fmt"
	"net/http"

	"github.com/golang/glog"
)

var authEmail *user.Email

type SignupController struct {
	controllers.BaseController
}

func (s *SignupController) Get() {

}

// Signup ..
// @router /user/signup [post]
func (s *SignupController) Signup() {
	username := s.GetString("username")
	password := s.GetString("password")
	email := s.GetString("email")

	data := map[string]interface{}{}
	dbUser := user.User{}
	authEmail = new(user.Email)

	dbUser.Name = username
	dbUser.Password = password
	dbUser.Email = email

	err := dbUser.Signup(&dbUser)
	if err != nil {
		if err.Error() == types.UsernameExErr {
			data["info"] = err.Error()
			s.ServerOk(data)
			glog.Infof("signup info[%s]", err.Error())
		} else {
			data["error"] = err.Error()
			s.ServerError(data, http.StatusInternalServerError)
			glog.Errorf("signup error[%s]", err.Error())
		}
		return
	}

	authURL := authEmail.GenerateAuthURL(email)

	authEmail.SetTheme("用户帐号激活") //设置主题
	authEmail.SetEmailContent(authURL)

	err = authEmail.InitSendCfg(email, username)
	if err != nil {
		glog.Errorf("send emial init config error: [%s]", err.Error())
		data["error"] = err.Error()
		s.ServerError(data, http.StatusInternalServerError)
		return
	}

	err = authEmail.SendEmail()
	if err != nil {
		data["error"] = err.Error()
		s.ServerError(data, http.StatusInternalServerError)
		return
	}

	data["info"] = "注册成功,已发送激活邮件，请激活后登录！"
	glog.Infoln(data["info"])
	s.ServerOk(data)

	return
}

// ActiveAccount activation user account by check email
// @router /user/auth/?:token [get]
func (s *SignupController) ActiveAccount() {
	token := s.GetString("token")
	fmt.Println("token: " + token)
	data := map[string]interface{}{}
	if authEmail != nil {
		isAccess, email := authEmail.CheckEmailURL(token)

		if isAccess {
			dbUser := user.User{}
			err := dbUser.ActiveUserByEmail(email)
			if err != nil {
				glog.Errorf("active user by email error[%s]\n", err.Error())
				data["error"] = err.Error()
				s.ServerError(data, http.StatusInternalServerError)
				return
			}
		}
		data["info"] = fmt.Sprintf("Auth email[%s] status[%t]", email, isAccess)
		glog.Infof(data["info"].(string))
		s.ServerOk(data)
		authEmail = nil
		return
	}
	err := fmt.Errorf("AuthEmail always nil")
	data["error"] = fmt.Sprintf("Auth emial url error[%s]", err.Error())
	s.ServerError(data, http.StatusBadRequest)
	return

}

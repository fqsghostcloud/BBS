package user

import (
	"bbs/controllers"
	"bbs/models/types"
	"bbs/models/user"
	"net/http"

	"github.com/golang/glog"
)

type SignupController struct {
	controllers.BaseController
}

func (s *SignupController) Get() {

}

func (s *SignupController) Post() {
	username := s.GetString("username")
	password := s.GetString("password")
	email := s.GetString("email")

	data := map[string]interface{}{}
	dbUser := user.User{}
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

	data["info"] = "注册成功!"
	glog.Infoln(data["info"])
	s.ServerOk(data)

	return
}

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

	data := map[string]interface{}{}
	dbUser := user.User{}
	ok, err := dbUser.Signup(username, password)
	if err != nil {
		if err.Error() == types.UsernameExErr {
			data["info"] = err.Error()
			s.serverOk(data)
			glog.Infof("signup info[%s]", err.Error())
		} else {
			data["error"] = err.Error()
			s.serverError(data, http.StatusInternalServerError)
			glog.Errorf("signup error[%s]", err.Error())
		}
		return
	}

	if ok {
		data["info"] = "注册成功!"
		glog.Infoln(data["info"])
		s.serverOk(data)
	}
	return
}

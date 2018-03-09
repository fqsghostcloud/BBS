package user

import (
	"bbs/models/types"
	"fmt"
	"time"

	"github.com/astaxie/beego"

	"github.com/golang/glog"

	"github.com/astaxie/beego/orm"     // check signup email
	_ "github.com/go-sql-driver/mysql" // import mysql driver
)

// Manager user api
type Manager interface {
	Auth(username, password string) (bool, error)
	AddUser(userInfo *User) error
	Delete(userInfo *User) error
	UpdateUser(userInfo *User) error
	GetUserByName(username string) (*User, error)
	GetUserByEmail(email string) (string, error) //get username
	ExsitUser(username string) bool
	ExsitEmail(email string) bool
	FuzzySearch(littleName string) (*[]User, error) //模糊查找
	Search(username string) (*User, error)          // 精确查找

	ActiveUser(username string) error
	DeactiveUser(username string) error
	IsActive(username string) (bool, error)
	ActiveUserByEmail(email string) error
	InactiveUserByEmail(email string) error

	AuthPassword(user *User, password string) bool
}

// User ..
type User struct {
	ID         int       `orm:"pk;auto"`
	Name       string    `orm:"unique;size(100)"`
	Password   string    `orm:"size(100)"` // 类型大小???????????????
	Active     bool      `orm:"default(false)"`
	Status     bool      `orm:"default(false)"`
	Email      string    `orm:"size(100)"`
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"` // frist save,record time
}

func init() {
	// orm.RegisterDriver("mysql", orm.DRMySQL) mysql auto regist
	user := new(User)
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("mysqlurl"), 30) //set database url and time zone
	orm.RegisterModel(user)
	orm.RunSyncdb("default", false, false) // begin create table?????函数参数作用

	user.CreateDefaultUser()

}

var singleManager Manager

// NewManager .
func NewManager() Manager {
	if singleManager != nil {
		return singleManager
	}
	singleManager = new(User)
	return singleManager
}

// TableName define database table name
func (u *User) TableName() string {
	return "user"
}

// ExsitUser check user whether exsit
func (u *User) ExsitUser(username string) bool {
	o := orm.NewOrm()
	qs := o.QueryTable(u.TableName())
	return qs.Filter("Name", username).Exist()
}

// ExsitEmail check user email whether exsit
func (u *User) ExsitEmail(email string) bool {
	o := orm.NewOrm()
	qs := o.QueryTable(u.TableName())
	return qs.Filter("Email", email).Exist()
}

// GetUserByName get user by username
func (u *User) GetUserByName(username string) (*User, error) {
	var user User
	o := orm.NewOrm()
	qs := o.QueryTable(u.TableName())
	err := qs.Filter("Name", username).One(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// IsActive check user whether activve
func (u *User) IsActive(username string) (bool, error) {
	user := new(User)

	o := orm.NewOrm()
	qs := o.QueryTable(u.TableName())
	err := qs.Filter("Name", username).One(user)

	if err != nil {
		return false, err
	}

	return user.Active == true, nil
}

// AddUser .
func (u *User) AddUser(userInfo *User) error {
	if u.ExsitUser(userInfo.Name) {
		return fmt.Errorf("%s", types.UsernameExErr)
	}

	if u.ExsitEmail(userInfo.Email) {
		return fmt.Errorf("%s", "此邮箱已经注册")
	}

	o := orm.NewOrm()
	_, err := o.Insert(userInfo)
	if err != nil {
		return err
	}

	return nil
}

// Delete delete user ..
func (u *User) Delete(userInfo *User) error {
	if u.ExsitUser(userInfo.Name) {
		o := orm.NewOrm()
		_, err := o.Delete(userInfo)
		if err != nil {
			return err
		}
		return nil
	}

	return fmt.Errorf(types.UserNotExsit)
}

// UpdateUser ..
func (u *User) UpdateUser(userInfo *User) error {
	if u.ExsitUser(userInfo.Name) {
		return fmt.Errorf("%s", types.UsernameExErr)
	}

	o := orm.NewOrm()
	_, err := o.Update(userInfo)
	if err != nil {
		return err
	}

	return nil
}

// GetUserByEmail ..
func (u *User) GetUserByEmail(email string) (string, error) {
	o := orm.NewOrm()
	user := new(User)
	qs := o.QueryTable(u.TableName())
	err := qs.Filter("Email", email).One(user)
	if err != nil {
		return "", err
	}

	return user.Name, nil

}

// ActiveUserByEmail ..
func (u *User) ActiveUserByEmail(email string) error {
	username, err := u.GetUserByEmail(email)
	if err != nil {
		return err
	}

	err = u.ActiveUser(username)
	if err != nil {
		return err
	}

	return nil
}

// InactiveUserByEmail ..
func (u *User) InactiveUserByEmail(email string) error {
	username, err := u.GetUserByEmail(email)
	if err != nil {
		return err
	}

	err = u.DeactiveUser(username)
	if err != nil {
		return err
	}

	return nil
}

// ActiveUser active user
func (u *User) ActiveUser(username string) error {
	if !u.ExsitUser(username) {
		return fmt.Errorf("%s", types.UserNotExsit)
	}

	o := orm.NewOrm()

	isActive, err := u.IsActive(username)
	if err != nil {
		return err
	}

	if isActive {
		err := fmt.Errorf("user[%s] already active", username)
		glog.Infof(err.Error())

		return err
	}

	qs := o.QueryTable(u.TableName())
	_, err = qs.Filter("Name", username).Update(orm.Params{"Active": true})
	if err != nil {
		return err
	}
	glog.Infof("active user[%s] success\n", username)

	return nil
}

// DeactiveUser inactive user
func (u *User) DeactiveUser(username string) error {
	if !u.ExsitUser(username) {
		return fmt.Errorf("%s", types.UserNotExsit)
	}

	isActive, err := u.IsActive(username)
	if err != nil {
		return err
	}

	if !isActive {
		err := fmt.Errorf("user[%s] already inactive", username)
		return err
	}

	o := orm.NewOrm()
	qs := o.QueryTable(u.TableName())

	_, err = qs.Filter("Name", username).Update(orm.Params{"Active": false})
	if err != nil {
		return err
	}

	return nil
}

// AuthPassword ..
func (u *User) AuthPassword(user *User, password string) bool {
	return user.Password == password
}

// Auth is check username and auth password
func (u *User) Auth(username, password string) (bool, error) {
	user, err := u.GetUserByName(username)
	if err != nil {
		glog.Errorf("get by username error[%s]", err.Error())
		return false, err
	}

	IsActive, err := u.IsActive(username)
	if err != nil {
		glog.Errorf("auth user is active error[%s]", err.Error())
		return false, err
	}

	if IsActive {
		return u.AuthPassword(user, password), nil
	}

	return false, fmt.Errorf("%s", types.UserLogForbidden)

}

// FuzzySearch 模糊查找返回匹配的用户
func (u *User) FuzzySearch(littleName string) (*[]User, error) {
	user := []User{}

	o := orm.NewOrm()
	qs := o.QueryTable(u.TableName())
	_, err := qs.Filter("Name__icontains", littleName).All(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// Search 精准查找返回用户公开信息
func (u *User) Search(username string) (*User, error) {
	return u.GetUserByName(username)
}

// CreateDefaultUser create test account
func (u *User) CreateDefaultUser() {
	user := new(User)
	user.Name = "admin"
	user.Password = "root"
	user.Active = true

	o := orm.NewOrm()

	if created, _, err := o.ReadOrCreate(user, "Name", "Password", "Active"); err == nil {
		if created {
			fmt.Printf("\nalready create user[%s]\n", user.Name)
		} else {
			fmt.Printf("\ncreate user[%s] success!\n", user.Name)
		}
	}
}

package user

import (
	"bbs/models/types"
	"fmt"
	"time"

	"github.com/golang/glog"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import mysql driver
)

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
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(192.168.34.140)/bbs?charset=utf8&loc=Asia%2FShanghai", 30) //set database url and time zone
	orm.RegisterModel(user)
	orm.RunSyncdb("default", false, false) // begin create table?????函数参数作用

	user.CreateDefaultUser()

}

// TableName define database table name
func (u *User) TableName() string {
	return "user"
}

func (u *User) Signup(userInfo *User) error {
	if u.ExsitUser(userInfo.Name) {
		return fmt.Errorf("%s", types.UsernameExErr)
	}

	o := orm.NewOrm()
	_, err := o.Insert(userInfo)
	if err != nil {
		return err
	}

	return nil
}

func (u *User) DelUser(userInfo *User) error {
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

// Auth is check username and auth password
func (u *User) Auth(username, password string) (bool, error) {
	user, err := u.GetByUsername(username)
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

func (u *User) AuthPassword(user *User, password string) bool {
	return user.Password == password
}

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

// GetUserInfo 获取用户详细信息或者模糊信息(username,isDetail)
func (u *User) GetUserInfo(username string, isDetail bool) {

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
	return u.GetByUsername(username)
}

func (u *User) GetByUsername(username string) (*User, error) {
	var user User
	o := orm.NewOrm()
	qs := o.QueryTable(u.TableName())
	err := qs.Filter("Name", username).One(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *User) ExsitUser(username string) bool {
	o := orm.NewOrm()
	qs := o.QueryTable(u.TableName())
	return qs.Filter("Name", username).Exist()
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

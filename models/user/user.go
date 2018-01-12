package user

import (
	"fmt"
	"time"

	"github.com/golang/glog"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import mysql driver
)

type User struct {
	Id         int       `orm:"pk;auto"`
	Name       string    `orm:"unique;size(100)"`
	Password   string    `orm:"size(100)"` // 类型大小？
	Active     bool      `orm:"default(false)"`
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"` // frist save,record time
}

func init() {
	// orm.RegisterDriver("mysql", orm.DRMySQL) mysql auto regist
	user := new(User)
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(192.168.34.139)/bbs?charset=utf8&loc=Asia%2FShanghai", 30) //set database url and time zone
	orm.RegisterModel(user)
	orm.RunSyncdb("default", false, false) // begin create table

	user.CreateDefaultUser()

}

// TableName define database table name
func (u *User) TableName() string {
	return "user"
}

// Auth is check username and auth password
func (u *User) Auth(username, password string) (bool, error) {
	user, err := u.GetByUsername(username)
	if err != nil {
		glog.Errorf("get by username error[%s]", err.Error())
		return false, err
	}

	return u.AuthPassword(user, password), nil

}

func (u *User) AuthPassword(user *User, password string) bool {
	return user.Password == password
}

func (u *User) GetByUsername(username string) (*User, error) {
	var user User
	o := orm.NewOrm()
	qs := o.QueryTable("user")
	err := qs.Filter("Name", username).One(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *User) ExsitUser(username string) bool {
	o := orm.NewOrm()
	qs := o.QueryTable("user")
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

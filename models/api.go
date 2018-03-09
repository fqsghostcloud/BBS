package models

import (
	"bbs/models/user"
)

// Manager api for all models
type Manager interface {
	// user
	Login(username, password string) (bool, error)
	SignUp(userInfo *user.User) error
	UpdateUser(user *user.User) error
	ActiveUser(username string) error
	ActiveAccount(email string) error
	DeactiveUser(username string) error
	DelUser(user *user.User) error
	FuzzySearch(littleName string) (*[]user.User, error) //模糊查找
	Search(allname string) (*user.User, error)           //精准查找
}

type managerImpl struct {
	user user.Manager
}

var singleManager Manager

// NewManager .
func NewManager() Manager {
	if singleManager != nil {
		return singleManager
	}
	singleManager = &managerImpl{
		user: user.NewManager(),
	}
	return singleManager
}

// FuzzySearch 模糊查找
func (m *managerImpl) FuzzySearch(littleName string) (*[]user.User, error) {
	return m.user.FuzzySearch(littleName)
}

// Search 精准查找
func (m *managerImpl) Search(allName string) (*user.User, error) {
	return m.user.Search(allName)
}

// DelUser .
func (m *managerImpl) DelUser(user *user.User) error {
	return m.user.Delete(user)
}

// ActiveUser ..
func (m *managerImpl) ActiveUser(username string) error {
	return m.user.ActiveUser(username)
}

// DeactiveUser
func (m *managerImpl) DeactiveUser(username string) error {
	return m.user.DeactiveUser(username)
}

// Login
func (m *managerImpl) Login(username, password string) (bool, error) {
	return m.user.Auth(username, password)
}

// UpdateUser
func (m *managerImpl) UpdateUser(user *user.User) error {
	return m.user.UpdateUser(user)
}

// Sign
func (m *managerImpl) SignUp(user *user.User) error {
	return m.user.AddUser(user)
}

// ActiveAccount
func (m *managerImpl) ActiveAccount(email string) error {
	return m.user.ActiveUserByEmail(email)
}

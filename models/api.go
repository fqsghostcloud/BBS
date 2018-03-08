package models

import (
	"bbs/models/user"
)

// Manager api for all models
type Manager interface {
	// user
	FuzzySearch(littleName string) (*[]user.User, error)
	Search(username string) (*user.User, error)
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
	singleManager = new(managerImpl)
	return singleManager
}

// FuzzySearch 模糊查找
func (m *managerImpl) FuzzySearch(littleName string) (*[]user.User, error) {
	return m.user.FuzzySearch(littleName)
}

// Search 精准查找
func (m *managerImpl) Search(username string) (*user.User, error) {
	return m.user.Search(username)
}

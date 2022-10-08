package services

import (
	"clean_architecture/feature/user/domain"
	"errors"
	"strings"
)

type UserBussines struct {
	mdl domain.UserData
}

func New(ud domain.UserData) domain.UserService {
	return &UserBussines{
		mdl: ud,
	}
}

// AddUser implements domain.UserService
func (ub *UserBussines) AddUser(newUser domain.User) (domain.User, error) {
	res, err := ub.mdl.Insert(newUser)
	var customErorr error
	if err != nil {
		if strings.Contains(err.Error(), "database") {
			customErorr = errors.New("cant insert to database")
		}
		return domain.User{}, customErorr
	}
	return res, nil
}

// ShowAllUser implements domain.UserService
func (ub *UserBussines) ShowAllUser() ([]domain.User, error) {
	res, err := ub.mdl.GetAll()
	var customErorr error
	if err != nil {
		if strings.Contains(err.Error(), "database getall vendor") {
			customErorr = errors.New("cant getAll from database")
		}
		return nil, customErorr
	}
	return res, nil
}

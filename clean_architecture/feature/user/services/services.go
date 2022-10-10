package services

import (
	"clean_architecture/feature/user/domain"
	"errors"
	"strings"
)

type UserBussines struct {
	mdl domain.UserData
}

// LoginUser implements domain.UserService

func New(ud domain.UserData) domain.UserService {
	return &UserBussines{
		mdl: ud,
	}
}

func (ub *UserBussines) Login(Nama string, HP string) (domain.User, error) {
	res, err := ub.mdl.GetUser(Nama, HP)
	var customErorr error
	if err != nil {
		if strings.Contains(err.Error(), "found") {
			customErorr = errors.New("not found in database")
		}
		return domain.User{}, customErorr
	}
	return res, nil
}

// AddUser implements domain.UserService
func (ub *UserBussines) Register(newUser domain.User) (domain.User, error) {
	res, err := ub.mdl.Insert(newUser)
	//var customErorr error
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return domain.User{}, errors.New("rejected from database")
		}

		return domain.User{}, errors.New("some problem on database")
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

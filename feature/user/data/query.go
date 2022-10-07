package data

import (
	"clean_architecture/feature/user/domain"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type UserQuery struct {
	db *gorm.DB
}

func New(DB *gorm.DB) domain.UserData {
	return &UserQuery{
		db: DB,
	}
}

// GetAll implements domain.UserData
func (uq *UserQuery) GetAll() ([]domain.User, error) {
	var res []domain.User
	var resTmp []UserModel
	var err error
	err = uq.db.Find(&resTmp).Error
	if err != nil {
		log.Error("user model eror:", err.Error())
		return nil, err
	}

	res = ToDomainArray(resTmp)
	return res, nil

}

// Insert implements domain.UserData
func (uq *UserQuery) Insert(newUser domain.User) (domain.User, error) {
	var input UserModel
	input = FromDomain(newUser)
	if err := uq.db.Create(&input).Error; err != nil {
		log.Error("user model error:", err.Error())
		return domain.User{}, err
	}
	var res domain.User
	res = ToDomain(input)
	return res, nil
}

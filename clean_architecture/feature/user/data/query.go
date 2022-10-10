package data

import (
	"clean_architecture/feature/user/domain"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type UserQuery struct {
	db *gorm.DB
}

func NewU(DB *gorm.DB) domain.UserData {
	return &UserQuery{
		db: DB,
	}
}

// Login implements domain.UserData
func (uq *UserQuery) GetUser(Nama string, HP string) (domain.User, error) {
	var res domain.User
	var resQry UserModel
	if err := uq.db.First(&resQry, "nama =? and hp=?", resQry.Nama, resQry.HP).Error; err != nil {
		log.Error("user model eror:", err.Error())
		return domain.User{}, err
	}

	res = ToDomain(resQry)
	return res, nil
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
	//var cnv UserModel //ambil dari db
	cnv := FromDomain(newUser)
	if err := uq.db.Create(&cnv).Error; err != nil {
		log.Error("user model error:", err.Error())
		return domain.User{}, err
	}
	//var res domain.User
	newUser = ToDomain(cnv)
	return newUser, nil
}

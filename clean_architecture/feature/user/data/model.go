package data

import (
	"clean_architecture/feature/user/domain"

	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	Nama string `json:"nama" form:"nama"`
	HP   string `json:"hp" form:"hp"`
}

func FromDomain(du domain.User) UserModel {
	return UserModel{
		Model: gorm.Model{ID: du.ID},
		Nama:  du.Nama,
		HP:    du.HP,
	}
}

func ToDomain(um UserModel) domain.User {
	return domain.User{
		ID:   um.ID,
		Nama: um.Nama,
		HP:   um.HP,
	}
}

func ToDomainArray(au []UserModel) []domain.User {
	var res []domain.User
	for _, val := range au {
		res = append(res, domain.User{ID: val.ID, Nama: val.Nama, HP: val.HP})
	}
	return res
}

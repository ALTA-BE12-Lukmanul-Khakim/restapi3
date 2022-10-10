package data

import (
	"clean_architecture/feature/logistic/domain"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type LogisticQuery struct {
	db *gorm.DB
}

func NewV(DB *gorm.DB) domain.LogisticData {
	return &LogisticQuery{
		db: DB,
	}
}

func (vq *LogisticQuery) Insert(newLogistic domain.Logistic) (domain.Logistic, error) {
	var input LogisticModel
	input = FromDomain(newLogistic)
	if err := vq.db.Create(&input).Error; err != nil {
		log.Error("Logistic model error", err.Error())
		return domain.Logistic{}, err
	}
	var res domain.Logistic
	res = ToDomain(input)
	return res, nil
}

func (vq *LogisticQuery) GetAll() ([]domain.Logistic, error) {
	var res []domain.Logistic
	var resTmp []LogisticModel
	var err error

	if err := vq.db.Find(&resTmp).Error; err != nil {
		log.Error("Get all Logistic error", err.Error())
		return nil, err
	}
	res = ToDomainArray(resTmp)
	return res, err
}

func (vq *LogisticQuery) GetVen(expedisi string) ([]domain.Logistic, error) {
	var res []domain.Logistic
	var resTmp []LogisticModel
	var err error
	if expedisi == "" {
		if err := vq.db.Find(&resTmp, "Expedisi =?", expedisi).Error; err != nil {
			log.Error("Get Expedisi Logistic Error", err.Error())
			return nil, err
		}
	} else {
		if err := vq.db.Find(&resTmp).Error; err != nil {
			log.Error("Get Expedisi Logistic Error", err.Error())
			return nil, err
		}

	}
	res = ToDomainArray(resTmp)
	return res, err
}

package services

import (
	"clean_architecture/feature/logistic/domain"
	"errors"
	"strings"

	"github.com/labstack/gommon/log"
)

type LogisticBussines struct {
	mdv domain.LogisticData
}

func New(vd domain.LogisticData) domain.LogisticService {
	return &LogisticBussines{
		mdv: vd,
	}
}

func (vb *LogisticBussines) AddLogistic(newLogistic domain.Logistic) (domain.Logistic, error) {
	res, err := vb.mdv.Insert(newLogistic)
	var customError error
	if err != nil {
		log.Error(err.Error())
		if strings.Contains(err.Error(), "database Insert") {
			customError = errors.New("cant insert to database")
		}
		return domain.Logistic{}, customError
	}
	return res, nil
}

func (vb *LogisticBussines) ShowAllLogistic() ([]domain.Logistic, error) {
	res, err := vb.mdv.GetAll()
	var customError error
	if err != nil {
		log.Error(err.Error())
		if strings.Contains(err.Error(), "database getAll") {
			customError = errors.New("cant getAll from database")
		}
		return nil, customError
	}
	if len(res) == 0 {
		log.Info("no data")
		return nil, errors.New("no data")
	}
	return res, nil
}

func (vb *LogisticBussines) ShowExpLogistic(expedisi string) ([]domain.Logistic, error) {
	res, err := vb.mdv.GetVen(expedisi)
	if err != nil {
		log.Error(err.Error())
		if strings.Contains(err.Error(), "database ") {
			return nil, errors.New("database error")
		} else if strings.Contains(err.Error(), "found ") {
			return nil, errors.New("no data")
		}
	}
	return res, nil
}

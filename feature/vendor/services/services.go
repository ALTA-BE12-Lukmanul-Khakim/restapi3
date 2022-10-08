package services

import (
	"clean_architecture/feature/vendor/domain"
	"errors"
	"strings"

	"github.com/labstack/gommon/log"
)

type VendorBussines struct {
	mdv domain.VendorData
}

func New(vd domain.VendorData) domain.VendorService {
	return &VendorBussines{
		mdv: vd,
	}
}

func (vb *VendorBussines) AddVendor(newVendor domain.Vendor) (domain.Vendor, error) {
	res, err := vb.mdv.Insert(newVendor)
	var customError error
	if err != nil {
		log.Error(err.Error())
		if strings.Contains(err.Error(), "database Insert") {
			customError = errors.New("cant insert to database")
		}
		return domain.Vendor{}, customError
	}
	return res, nil
}

func (vb *VendorBussines) ShowAllVendor() ([]domain.Vendor, error) {
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

func (vb *VendorBussines) ShowExpVendor(expedisi string) ([]domain.Vendor, error) {
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

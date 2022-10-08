package data

import (
	"clean_architecture/feature/vendor/domain"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type VendorQuery struct {
	db *gorm.DB
}

func New(DB *gorm.DB) domain.VendorData {
	return &VendorQuery{
		db: DB,
	}
}

func (vq *VendorQuery) Insert(newVendor domain.Vendor) (domain.Vendor, error) {
	var input VendorModel
	input = FromDomain(newVendor)
	if err := vq.db.Create(&input).Error; err != nil {
		log.Error("vendor model error", err.Error())
		return domain.Vendor{}, err
	}
	var res domain.Vendor
	res = ToDomain(input)
	return res, nil
}

func (vq *VendorQuery) GetAll() ([]domain.Vendor, error) {
	var res []domain.Vendor
	var resTmp []VendorModel
	var err error

	if err := vq.db.Find(&resTmp).Error; err != nil {
		log.Error("Get all vendor error", err.Error())
		return nil, err
	}
	res = ToDomainArray(resTmp)
	return res, err
}

func (vq *VendorQuery) GetVen(expedisi string) ([]domain.Vendor, error) {
	var res []domain.Vendor
	var resTmp []VendorModel
	var err error

	if err := vq.db.Find(&resTmp).Where(&resTmp, "Expedisi =?", expedisi).Error; err != nil {
		log.Error("Get Expedisi Vendor Error", err.Error())
		return nil, err
	}
	res = ToDomainArray(resTmp)
	return res, err
}

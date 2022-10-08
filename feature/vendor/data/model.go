package data

import (
	"clean_architecture/feature/vendor/domain"

	"gorm.io/gorm"
)

type VendorModel struct {
	gorm.Model
	NamaVen  string
	Expedisi string
}

func FromDomain(dv domain.Vendor) VendorModel {
	return VendorModel{
		NamaVen:  dv.NamaVen,
		Expedisi: dv.Expedisi,
	}
}

func ToDomain(vm VendorModel) domain.Vendor {
	return domain.Vendor{
		NamaVen:  vm.NamaVen,
		Expedisi: vm.Expedisi,
	}
}

func ToDomainArray(av []VendorModel) []domain.Vendor {
	var res []domain.Vendor
	for _, val := range av {
		res = append(res, domain.Vendor{ID: val.ID, NamaVen: val.NamaVen, Expedisi: val.Expedisi})
	}
	return res
}
